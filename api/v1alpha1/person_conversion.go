package v1alpha1

import (
	"strings"

	"github.com/simingweng/multi-ver-sample/api/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

var _ conversion.Convertible = &Person{}

func (src *Person) ConvertTo(dst conversion.Hub) error {
	dstP := dst.(*v1beta1.Person)
	personlog.Info("convert", "from", src.APIVersion, "to", dstP.APIVersion)

	ss := strings.Split(src.Spec.Name, ",")
	dstP.Spec.FirstName = ss[0]
	dstP.Spec.LastName = ss[1]

	dstP.ObjectMeta = src.ObjectMeta

	return nil
}

// ConvertFrom implements conversion.Convertible
func (dst *Person) ConvertFrom(src conversion.Hub) error {
	srcP := src.(*v1beta1.Person)
	personlog.Info("convert", "from", srcP.APIVersion, "to", dst.APIVersion)

	dst.Spec.Name = srcP.Spec.FirstName + "," + srcP.Spec.LastName

	dst.ObjectMeta = srcP.ObjectMeta

	return nil
}
