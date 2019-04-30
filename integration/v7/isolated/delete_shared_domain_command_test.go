package isolated

import (
	"regexp"

	"code.cloudfoundry.org/cli/integration/helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("delete-shared-domain command", func() {
	When("--help flag is given", func() {
		It("Displays command usage to output", func() {
			session := helpers.CF("delete-shared-domain", "--help")

			Eventually(session).Should(Say("NAME:"))
			Eventually(session).Should(Say(`\s+delete-shared-domain - Delete a shared domain`))
			Eventually(session).Should(Say("USAGE:"))
			Eventually(session).Should(Say(`\s+cf delete-shared-domain DOMAIN \[-f\]`))
			Eventually(session).Should(Say("OPTIONS:"))
			Eventually(session).Should(Say(`\s+-f\s+Force deletion without confirmation`))
			Eventually(session).Should(Say("SEE ALSO:"))
			Eventually(session).Should(Say(`\s+delete-domain, domains`))
			Eventually(session).Should(Exit(0))
		})
	})

	When("the environment is set up correctly", func() {
		var (
			buffer     *Buffer
			orgName    string
			spaceName  string
			domainName string
			username   string
		)

		BeforeEach(func() {
			buffer = NewBuffer()
			domainName = helpers.NewDomainName()
			orgName = helpers.NewOrgName()
			spaceName = helpers.NewSpaceName()

			username, _ = helpers.GetCredentials()
			helpers.SetupCF(orgName, spaceName)

			session := helpers.CF("create-shared-domain", domainName)
			Eventually(session).Should(Exit(0))
		})

		When("the -f flag is not given", func() {
			When("the user enters 'y'", func() {
				BeforeEach(func() {
					_, err := buffer.Write([]byte("y\n"))
					Expect(err).ToNot(HaveOccurred())
				})

				It("it asks for confirmation and deletes the domain", func() {
					session := helpers.CFWithStdin(buffer, "delete-shared-domain", domainName)
					Eventually(session).Should(Say(regexp.QuoteMeta(`Deleting domain %s as %s...`), domainName, username))
					Eventually(session).Should(Say("\n\nThis action impacts all orgs using this domain."))
					Eventually(session).Should(Say("Deleting it will remove associated routes and could make any app with this domain, in any org, unreachable."))
					Eventually(session).Should(Say(`Are you sure you want to delete the domain %s\?`, domainName))
					Eventually(session).Should(Say("OK"))
					Eventually(session).Should(Exit(0))
				})
			})

			When("the user enters 'n'", func() {
				BeforeEach(func() {
					_, err := buffer.Write([]byte("n\n"))
					Expect(err).ToNot(HaveOccurred())
				})

				It("it asks for confirmation and does not delete the domain", func() {
					session := helpers.CFWithStdin(buffer, "delete-shared-domain", domainName)
					Eventually(session).Should(Say(regexp.QuoteMeta(`Deleting domain %s as %s...`), domainName, username))
					Eventually(session).Should(Say("\n\nThis action impacts all orgs using this domain."))
					Eventually(session).Should(Say("Deleting it will remove associated routes and could make any app with this domain, in any org, unreachable."))
					Eventually(session).Should(Say(`Are you sure you want to delete the domain %s\?`, domainName))
					Consistently(session).ShouldNot(Say("OK"))
					Eventually(session).Should(Exit(0))
				})
			})
		})

		When("the -f flag is given", func() {
			It("it deletes the domain without asking for confirmation", func() {
				session := helpers.CFWithStdin(buffer, "delete-shared-domain", domainName, "-f")
				Eventually(session).Should(Say(regexp.QuoteMeta(`Deleting domain %s as %s...`), domainName, username))
				Consistently(session).ShouldNot(Say("Are you sure"))
				Eventually(session).Should(Say("OK"))
				Eventually(session).Should(Exit(0))
			})
		})
	})
})
