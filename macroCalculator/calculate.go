package macroCalculator


import(

	parser "foodApp/tools/parser"
	"log"
)

func CollectProfileInfo() {
	profiles := parser.ReadProfileJSON("./data/profiles.json")
	
	for _, p := range profiles {
		log.Print(p.CalculateBMR())
	}
}