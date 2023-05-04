package persistence

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/arturiamu/lplms-public_cloud/common"
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/arturiamu/lplms-public_cloud/utils/filter"
	"github.com/arturiamu/lplms-public_cloud/utils/pointerutil"
	"golang.org/x/crypto/ssh"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubevirtv1 "kubevirt.io/api/core/v1"
	"log"
	"strings"
)

// CreateKeypair
// 调用CreateKeyPair创建一对SSH密钥对或导入公钥。如果不传入公钥信息，则表示生成新密钥对，我们会为您保管密钥的公钥部分，
// 并返回未加密的PEM编码的PKCS#8格式私钥。您需要自行妥善保管私钥部分。
// 您在每个地域的密钥对数最高为500对。更多详情，请参见使用限制。
func (c *ComputeRepo) CreateKeypair(args *entity.KeypairCreateArg) (*entity.KeypairCreateResp, error) {
	var (
		ns              = args.ProjectID
		k               = c.k8Virt.CoreV1().Secrets(ns)
		pubKey, prvtKey []byte
	)

	if args.PublicKeyBody == nil {
		var err error
		pubKey, prvtKey, err = generateSSHKey()
		if err != nil {
			return nil, err
		}
	} else {
		pubKey = []byte(*args.PublicKeyBody)
	}

	fp, err := fingerPrint(pubKey)
	if err != nil {
		return nil, err
	}

	s, err := k.Create(context.Background(), &corev1.Secret{
		ObjectMeta: v1.ObjectMeta{
			Name: args.KeyPairName,
			Annotations: map[string]string{
				common.AnnotationFingerPrint: fp,
			},
		},
		Type: corev1.SecretTypeOpaque,
		Data: map[string][]byte{"key1": pubKey},
	}, v1.CreateOptions{})

	if err != nil {
		return nil, err
	}

	return &entity.KeypairCreateResp{
		KeyPairName:    s.ObjectMeta.Name,
		PrivateKeyBody: pointerutil.Pointer(string(prvtKey)),
		FingerPrint:    fp,
	}, nil
}

func (c *ComputeRepo) DeleteKeypair(args *entity.KeypairDeleteArg) (*entity.KeypairDeleteResp, error) {
	err := c.k8Virt.CoreV1().Secrets(args.ProjectID).Delete(context.Background(), args.KeyPairName, v1.DeleteOptions{})
	return nil, err
}

func (c *ComputeRepo) UpdateKeypair(args *entity.KeypairUpdateArg) (*entity.KeypairUpdateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) GetKeypair(args *entity.KeypairGetArg) (*entity.KeypairGetResp, error) {
	//TODO implement me
	panic("implement me")
}

// ListKeypair
// 调用DescribeKeyPairs查询一个或多个密钥对。
func (c *ComputeRepo) ListKeypair(args *entity.KeypairListArg) (*entity.KeypairListResp, error) {
	resp, err := c.k8Virt.CoreV1().Secrets(args.ProjectID).List(context.Background(), v1.ListOptions{})
	if err != nil {
		return nil, err
	}

	res := &entity.KeypairListResp{
		KeyPairs: make([]entity.Keypair, 0, len(resp.Items)),
	}

	if args.KeyPairName != nil {
		for _, v := range resp.Items {
			if v.Name == *args.KeyPairName {
				res.KeyPairs = append(res.KeyPairs, entity.Keypair{
					KeyPairName: v.Name,
					FingerPrint: v.Annotations[common.AnnotationFingerPrint],
				})
			}
		}
		return res, nil
	}

	for _, v := range resp.Items {
		if strings.HasPrefix(v.Name, "default-token-") {
			continue
		}

		res.KeyPairs = append(res.KeyPairs, entity.Keypair{
			KeyPairName: v.Name,
			FingerPrint: v.Annotations[common.AnnotationFingerPrint],
		})
	}

	return res, nil
}

// AttachKeyPair
// 调用AttachKeyPair绑定一个SSH密钥对到一台Linux实例。
// 当您使用该接口时，需要注意：
// Windows实例不支持SSH密钥对。
// 绑定SSH密钥对后，将禁用用户名加密码的验证方式。
// 如果实例处于**运行中**（`Running`）状态，重启实例（RebootServer）后，SSH密钥对生效。
// 如果实例处于**已停止**（`Stopped`）状态，启动实例（StartServer）后，SSH密钥对生效。
// 如果实例已经绑定了SSH密钥对，新的SSH密钥对自动替换原来的SSH密钥对。
// 调用DeleteKeyPairs删除一对SSH密钥对。
func (c *ComputeRepo) AttachKeyPair(args *entity.KeypairAttachArg) (*entity.KeypairAttachResp, error) {
	var (
		ns = args.ProjectID
		k  = c.k8Virt.VirtualMachine(ns)
	)

	oldServer, err := k.Get(args.ServerID, &v1.GetOptions{})
	if err != nil {
		return nil, err
	}

	newServer := oldServer.DeepCopy()

	accessCredential := kubevirtv1.AccessCredential{}
	accessCredential.SSHPublicKey = &kubevirtv1.SSHPublicKeyAccessCredential{
		Source: kubevirtv1.SSHPublicKeyAccessCredentialSource{
			Secret: &kubevirtv1.AccessCredentialSecretSource{
				SecretName: args.KeyPairName,
			},
		},
		PropagationMethod: kubevirtv1.SSHPublicKeyAccessCredentialPropagationMethod{
			QemuGuestAgent: &kubevirtv1.QemuGuestAgentSSHPublicKeyAccessCredentialPropagation{
				Users: []string{"root"},
			},
		},
	}

	newServer.Spec.Template.Spec.AccessCredentials = append(newServer.Spec.Template.Spec.AccessCredentials, accessCredential)

	_, err = k.Update(newServer)

	return nil, err
}

// DetachKeyPair
// 删除SSH密钥对后，您需要注意：
// - 无法通过GetKeyPairs查询该SSH密钥对。
// - 若已有ECS实例绑定了该SSH密钥对：
// - LPLMS不再为您保存该SSH密钥对，但是实例可以正常使用该SSH密钥对。
// - 查询实例信息时（GetServers），会显示SSH密钥对名称（KeyPairName），但不再显示其他相关信息。
func (c *ComputeRepo) DetachKeyPair(args *entity.KeypairDetachArg) (*entity.KeypairDetachResp, error) {
	var (
		ns = args.ProjectID
		k  = c.k8Virt.VirtualMachine(ns)
	)

	oldServer, err := k.Get(args.ServerID, &v1.GetOptions{})
	if err != nil {
		return nil, err
	}

	newServer := oldServer.DeepCopy()

	result := filter.Filter(
		newServer.Spec.Template.Spec.AccessCredentials,
		func(item kubevirtv1.AccessCredential, idx int) bool {
			if args.KeyPairName != item.SSHPublicKey.Source.Secret.SecretName {
				return true
			}

			return false
		},
	)

	newServer.Spec.Template.Spec.AccessCredentials = result

	_, err = k.Update(newServer)
	return nil, err
}

///////////////////////////// help functions /////////////

func (c *ComputeRepo) getPubKey(keypairName string, ns string) (pubKey string, err error) {
	resp, err := c.k8Virt.CoreV1().Secrets(ns).Get(context.Background(), keypairName, v1.GetOptions{})
	if err != nil {
		return
	}

	if resp != nil {
		pubKey = string(resp.Data["key1"])
	}

	return
}

func generateSSHKey() (pubKey, prvtKey []byte, err error) {
	privateKey, err := generatePrivateKey(2048)
	if err != nil {
		return
	}

	pubKey, err = generatePublicKey(&privateKey.PublicKey)
	if err != nil {
		return
	}

	prvtKey = encodePrivateKeyToPEM(privateKey)
	return
}

// encodePrivateKeyToPEM
// encodes Private Key from RSA to PEM format
func encodePrivateKeyToPEM(privateKey *rsa.PrivateKey) []byte {
	// Get ASN.1 DER format
	privDER := x509.MarshalPKCS1PrivateKey(privateKey)

	// pem.Block
	privBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privDER,
	}

	// Private key in PEM format
	privatePEM := pem.EncodeToMemory(&privBlock)

	return privatePEM
}

// generatePublicKey
// take a rsa.PublicKey and return bytes suitable for writing to .pub file
// returns in the format "ssh-rsa ..."
func generatePublicKey(privatekey *rsa.PublicKey) ([]byte, error) {
	publicRsaKey, err := ssh.NewPublicKey(privatekey)
	if err != nil {
		log.Println("err:", err)
		return nil, err
	}

	pubKeyBytes := ssh.MarshalAuthorizedKey(publicRsaKey)

	log.Println("Public key generated")
	return pubKeyBytes, nil
}

// generatePrivateKey
// creates a RSA Private Key of specified byte size
func generatePrivateKey(bitSize int) (*rsa.PrivateKey, error) {
	// Private Key generation
	privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return nil, err
	}

	// Validate Private Key
	err = privateKey.Validate()
	if err != nil {
		return nil, err
	}

	log.Println("Private Key generated")
	return privateKey, nil
}

func fingerPrint(pub []byte) (fp string, err error) {
	key, _, _, _, err := ssh.ParseAuthorizedKey(pub)
	if err != nil {
		return
	}

	fp = ssh.FingerprintLegacyMD5(key)
	return
}
