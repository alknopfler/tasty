/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"strings"
	"tasty/pkg/utils"

	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search matching operators",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var operator string
		var currentoperator string
		if len(args) != 1 {
			log.Printf("Usage: tasty search OPERATOR_NAME")
		} else {
			operator = args[0]
		}
		dynamic := utils.GetDynamicClient()
		packagemanifests := schema.GroupVersionResource{Group: "packages.operators.coreos.com", Version: "v1", Resource: "packagemanifests"}
		list, err := dynamic.Resource(packagemanifests).Namespace("openshift-marketplace").List(context.TODO(), metav1.ListOptions{})
		utils.Check(err)
		for _, d := range list.Items {
			currentoperator = d.GetName()
			if strings.Contains(currentoperator, operator) {
				fmt.Println(currentoperator)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
