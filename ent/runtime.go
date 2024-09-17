// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/dipper-iot/dipper-engine-server/ent/rulechan"
	"github.com/dipper-iot/dipper-engine-server/ent/rulenode"
	"github.com/dipper-iot/dipper-engine-server/ent/schema"
	"github.com/dipper-iot/dipper-engine-server/ent/session"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	rulechanFields := schema.RuleChan{}.Fields()
	_ = rulechanFields
	// rulechanDescName is the schema descriptor for name field.
	rulechanDescName := rulechanFields[1].Descriptor()
	// rulechan.NameValidator is a validator for the "name" field. It is called by the builders before save.
	rulechan.NameValidator = func() func(string) error {
		validators := rulechanDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// rulechanDescDescription is the schema descriptor for description field.
	rulechanDescDescription := rulechanFields[2].Descriptor()
	// rulechan.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	rulechan.DescriptionValidator = rulechanDescDescription.Validators[0].(func(string) error)
	// rulechanDescRootNode is the schema descriptor for root_node field.
	rulechanDescRootNode := rulechanFields[3].Descriptor()
	// rulechan.RootNodeValidator is a validator for the "root_node" field. It is called by the builders before save.
	rulechan.RootNodeValidator = func() func(string) error {
		validators := rulechanDescRootNode.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(root_node string) error {
			for _, fn := range fns {
				if err := fn(root_node); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// rulechanDescCreatedAt is the schema descriptor for created_at field.
	rulechanDescCreatedAt := rulechanFields[5].Descriptor()
	// rulechan.DefaultCreatedAt holds the default value on creation for the created_at field.
	rulechan.DefaultCreatedAt = rulechanDescCreatedAt.Default.(func() time.Time)
	// rulechanDescUpdatedAt is the schema descriptor for updated_at field.
	rulechanDescUpdatedAt := rulechanFields[6].Descriptor()
	// rulechan.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	rulechan.DefaultUpdatedAt = rulechanDescUpdatedAt.Default.(func() time.Time)
	rulenodeFields := schema.RuleNode{}.Fields()
	_ = rulenodeFields
	// rulenodeDescNodeID is the schema descriptor for node_id field.
	rulenodeDescNodeID := rulenodeFields[2].Descriptor()
	// rulenode.NodeIDValidator is a validator for the "node_id" field. It is called by the builders before save.
	rulenode.NodeIDValidator = func() func(string) error {
		validators := rulenodeDescNodeID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(node_id string) error {
			for _, fn := range fns {
				if err := fn(node_id); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// rulenodeDescRuleID is the schema descriptor for rule_id field.
	rulenodeDescRuleID := rulenodeFields[3].Descriptor()
	// rulenode.RuleIDValidator is a validator for the "rule_id" field. It is called by the builders before save.
	rulenode.RuleIDValidator = func() func(string) error {
		validators := rulenodeDescRuleID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(rule_id string) error {
			for _, fn := range fns {
				if err := fn(rule_id); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// rulenodeDescOption is the schema descriptor for option field.
	rulenodeDescOption := rulenodeFields[4].Descriptor()
	// rulenode.DefaultOption holds the default value on creation for the option field.
	rulenode.DefaultOption = rulenodeDescOption.Default.(map[string]interface{})
	// rulenodeDescDebug is the schema descriptor for debug field.
	rulenodeDescDebug := rulenodeFields[5].Descriptor()
	// rulenode.DefaultDebug holds the default value on creation for the debug field.
	rulenode.DefaultDebug = rulenodeDescDebug.Default.(bool)
	// rulenodeDescEnd is the schema descriptor for end field.
	rulenodeDescEnd := rulenodeFields[6].Descriptor()
	// rulenode.DefaultEnd holds the default value on creation for the end field.
	rulenode.DefaultEnd = rulenodeDescEnd.Default.(bool)
	// rulenodeDescCreatedAt is the schema descriptor for created_at field.
	rulenodeDescCreatedAt := rulenodeFields[7].Descriptor()
	// rulenode.DefaultCreatedAt holds the default value on creation for the created_at field.
	rulenode.DefaultCreatedAt = rulenodeDescCreatedAt.Default.(func() time.Time)
	// rulenodeDescUpdatedAt is the schema descriptor for updated_at field.
	rulenodeDescUpdatedAt := rulenodeFields[8].Descriptor()
	// rulenode.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	rulenode.DefaultUpdatedAt = rulenodeDescUpdatedAt.Default.(func() time.Time)
	sessionFields := schema.Session{}.Fields()
	_ = sessionFields
	// sessionDescIsTest is the schema descriptor for is_test field.
	sessionDescIsTest := sessionFields[2].Descriptor()
	// session.DefaultIsTest holds the default value on creation for the is_test field.
	session.DefaultIsTest = sessionDescIsTest.Default.(bool)
	// sessionDescInfinite is the schema descriptor for infinite field.
	sessionDescInfinite := sessionFields[3].Descriptor()
	// session.DefaultInfinite holds the default value on creation for the infinite field.
	session.DefaultInfinite = sessionDescInfinite.Default.(bool)
	// sessionDescData is the schema descriptor for data field.
	sessionDescData := sessionFields[4].Descriptor()
	// session.DefaultData holds the default value on creation for the data field.
	session.DefaultData = sessionDescData.Default.(map[string]interface{})
	// sessionDescResult is the schema descriptor for result field.
	sessionDescResult := sessionFields[5].Descriptor()
	// session.DefaultResult holds the default value on creation for the result field.
	session.DefaultResult = sessionDescResult.Default.(map[string]interface{})
	// sessionDescEndCount is the schema descriptor for end_count field.
	sessionDescEndCount := sessionFields[6].Descriptor()
	// session.DefaultEndCount holds the default value on creation for the end_count field.
	session.DefaultEndCount = sessionDescEndCount.Default.(int)
	// sessionDescTimeout is the schema descriptor for timeout field.
	sessionDescTimeout := sessionFields[7].Descriptor()
	// session.DefaultTimeout holds the default value on creation for the timeout field.
	session.DefaultTimeout = sessionDescTimeout.Default.(int)
	// sessionDescCreatedAt is the schema descriptor for created_at field.
	sessionDescCreatedAt := sessionFields[8].Descriptor()
	// session.DefaultCreatedAt holds the default value on creation for the created_at field.
	session.DefaultCreatedAt = sessionDescCreatedAt.Default.(func() time.Time)
	// sessionDescUpdatedAt is the schema descriptor for updated_at field.
	sessionDescUpdatedAt := sessionFields[9].Descriptor()
	// session.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	session.DefaultUpdatedAt = sessionDescUpdatedAt.Default.(func() time.Time)
}
