// Wrapper functions for "terraform destroy"
package terraform

import (
	"log"

	"github.com/gruntwork-io/terratest/shell"
	"github.com/gruntwork-io/terratest/util"
)

// Call Terraform Destroy on the template at "templatePath" with the given "vars"
func Destroy(templatePath string, vars map[string]string, logger *log.Logger) error {
	logger.Println("Destroy Terraform changes in folder", templatePath)
	return shell.RunCommand(shell.Command { Command: "terraform", Args: util.FormatArgs(vars, "destroy", "-force", "-input=false"), WorkingDir: templatePath, Env: terraformDebugEnv }, logger)
}

// Call Terraform Destroy on the template at "templatePath" with the given "vars", and return the output as a string
func DestroyAndGetOutput(templatePath string, vars map[string]string, logger *log.Logger) (string, error) {
	logger.Println("Destroy Terraform changes in folder", templatePath)
	return shell.RunCommandAndGetOutput(shell.Command { Command: "terraform", Args: util.FormatArgs(vars, "destroy", "-force", "-input=false"), WorkingDir: templatePath, Env: terraformDebugEnv }, logger)
}
