package manage_indexes

import (
	"fmt"
	"strconv"
)

func VerfiyIndexDefinitionFromScenario(results []IndexDefinition, scenario VectorIndexScenario) bool {
	localIsValid := true
	for _, result := range results {
		if result.Name != scenario.Name {
			localIsValid = false
			fmt.Printf("Expected the index name " + scenario.Name + " but got " + result.Name + "\n")
		}

		for ii, expectedFields := range scenario.Fields {
			if result.LatestDefinition.Fields[ii].Type != expectedFields.Type {
				localIsValid = false
				fmt.Printf("Expected the type " + expectedFields.Type + " but got " + result.LatestDefinition.Fields[ii].Type + "\n")
			}

			if result.LatestDefinition.Fields[ii].Path != expectedFields.Path {
				localIsValid = false
				fmt.Printf("Expected the path " + expectedFields.Path + " but got " + result.LatestDefinition.Fields[ii].Path + "\n")
			}

			if expectedFields.Type == "vectorSearch" {
				if result.LatestDefinition.Fields[ii].NumDimensions != expectedFields.NumDimensions {
					localIsValid = false
					fmt.Printf("Expected num dimensions to be %v, but got %v\n", strconv.Itoa(expectedFields.NumDimensions), strconv.Itoa(result.LatestDefinition.Fields[ii].NumDimensions))
				}

				if result.LatestDefinition.Fields[ii].Similarity != expectedFields.Similarity {
					localIsValid = false
					fmt.Printf("Expected the similarity " + expectedFields.Similarity + " but got " + result.LatestDefinition.Fields[ii].Similarity + "\n")
				}
			}
		}
	}
	return localIsValid
}
