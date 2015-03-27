package swagger_test

import (
	"net/http"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSwagger(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Swagger Suite")
}

var _ = BeforeSuite(func() {
})

var _ = AfterSuite(func() {
})

var err error

var _ = Describe("Swagger", func() {
	BeforeEach(func() {
	})
	Describe("petstore", func() {
		Context("specs", func() {
			It("should generate", func() {
				s := swagger.NewSchema()
				s.Set(swagger.SWAGGER_FIELD_VERSION, swagger.SWAGGER_VERSION_2_0)
				info := swagger.NewSchema()
				{
					info.Set(swagger.INFO_FIELD_TITLE, "Swagger Petstore")
					info.Set(swagger.INFO_FIELD_VERSION, "1.0.0")
					info.Set(swagger.INFO_FIELD_DESCRIPTION, "This is a sample server Petstore server.  You can find out more about Swagger at <a href=\"http://swagger.io\">http://swagger.io</a> or on irc.freenode.net, #swagger.  For this sample, you can use the api key \"special-key\" to test the authorization filters")
					info.Set(swagger.INFO_FIELD_TERMS_OF_SERVICE, "http://helloreverb.com/terms/")
					contact := swagger.NewSchema()
					{
						contact.Set(swagger.CONTACT_FIELD_EMAIL, "apiteam@wordnik.com")
						info.Set(swagger.INFO_FIELD_CONTACT, contact)
					}
					license := swagger.NewSchema()
					{
						license.Set(swagger.LICENSE_FIELD_NAME, "Apache 2.0")
						license.Set(swagger.LICENSE_FIELD_URL, "http://www.apache.org/licenses/LICENSE-2.0.html")
						info.Set(swagger.INFO_FIELD_LICENSE, license)
					}
					s.Set(swagger.SWAGGER_FIELD_INFO, info)
				}
				paths := swagger.NewSchema()
				{
					pi := swagger.NewSchema() // -> /pet
					{
						o := swagger.NewSchema() // Post
						{
							o.ExtendSlice(swagger.OPERATION_FIELD_TAGS, "pet")
							o.Set(swagger.OPERATION_FIELD_SUMMARY, "Add a new pet to the store")
							o.Set(swagger.OPERATION_FIELD_DESCRIPTION, "")
							o.Set(swagger.OPERATION_FIELD_OPERATION_ID, "addPet")

							o.ExtendSlice(swagger.OPERATION_FIELD_CONSUMES, swagger.MIME_TYPE_APPLICATION_JSON)
							o.ExtendSlice(swagger.OPERATION_FIELD_CONSUMES, swagger.MIME_TYPE_APPLICATION_XML)

							o.ExtendSlice(swagger.OPERATION_FIELD_PRODUCES, swagger.MIME_TYPE_APPLICATION_JSON)
							o.ExtendSlice(swagger.OPERATION_FIELD_PRODUCES, swagger.MIME_TYPE_APPLICATION_XML)
							responses := swagger.NewSchema()
							{
								r := swagger.NewSchema()
								{
									r.Set(swagger.RESPONSE_FIELD_DESCRIPTION, "Invalid input")
									responses.Set(http.StatusNotFound, r)
								}
								o.Set(swagger.OPERATION_FIELD_RESPONSES, responses)
							}
							schema := swagger.NewSchema()
							{
								schema.Set(swagger.REFERENCE_FIELD_REF, "#/definitions/Pet")
							}
							parameter := swagger.NewSchema()
							{
								parameter.Set(swagger.PARAMETER_FIELD_SCHEMA, schema)
								parameter.Set(swagger.PARAMETER_FIELD_IN, swagger.PARAMETER_TYPE_BODY)
								parameter.Set(swagger.PARAMETER_FIELD_TYPE, swagger.PARAMETER_DATA_TYPE_STRING)
								parameter.Set(swagger.PARAMETER_FIELD_DESCRIPTION, "Pet object that needs to be added to the store")
								o.ExtendSlice(swagger.OPERATION_FIELD_PARAMETERS, parameter)
							}
							sr := swagger.NewSchema()
							{
								sr.ExtendSlice("petstore_auth", "write:pets")
								sr.ExtendSlice("petstore_auth", "read:pets")
								o.ExtendSlice(swagger.OPERATION_FIELD_SECURITY, sr)
							}
						}
						pi.Set(swagger.PATH_ITEM_FIELD_POST, o)
						paths.Set("/pet", pi)
					}
					s.Set(swagger.SWAGGER_FIELD_PATHS, paths)
				}
				s.Generate(true)
				Expect(s.HasErrors()).To(Equal(false))
			})
		})
	})
})
