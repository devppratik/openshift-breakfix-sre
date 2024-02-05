package main

import (
	"fmt"
	"openshift-breakfix/challenges"

	"github.com/manifoldco/promptui"
)

func DisplayPrompt() {
	prompt := promptui.Select{
		Label: "Select the component",
		Items: []string{"etcd", "operator"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)
	if result == "etcd" {
		challenges.InitEtcdMembersDown()
	} else {
		fmt.Println("Work Under Progress")
	}
}
