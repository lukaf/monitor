package main

import (
	"flag"
	"log"

	monitor "github.com/lukaf/monitor/m"
)

var (
	accessKey *string = flag.String("a", "", "AccessKey (optional, if not provided environmental variable or instance profile is used)")
	secretKey *string = flag.String("s", "", "SecretKey (optional, if not provided environmental variable or instance profile is used)")
	region    *string = flag.String("r", "eu-west-1", "Region")

	unit     *string = flag.String("u", "", "Unit of the metric.")
	name     *string = flag.String("n", "", "Name of the metric.")
	resource *string = flag.String("rs", "", "Select resource")

	autoScalingGroup *string = flag.String("asg", "", "Name of the auto scaling group (optional).")
	namespace        *string = flag.String("ns", "", "Namespace of the metric.")
)

func init() {
	flag.Parse()

	if _, ok := monitor.Resources[*resource]; !ok {
		log.Fatalf(`Unknown monitor "%s". Available: %s`, *resource, monitor.ListResources())
	}
}

func main() {
	cw, err := monitor.NewCW(*accessKey, *secretKey, *region)
	if err != nil {
		log.Fatalln(err)
	}

	value, defaultUnit, defaultName := monitor.GetResource(*resource)

	if *unit == "" {
		*unit = defaultUnit
	}

	if *name == "" {
		*name = defaultName
	}

	if _, err := cw.PutMetrics(value, *unit, *name, *namespace, *autoScalingGroup); err != nil {
		log.Fatalln(err)
	}
}
