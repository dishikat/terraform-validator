// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import "reflect"

const CloudDeployDeliveryPipelineAssetType string = "clouddeploy.googleapis.com/DeliveryPipeline"

func resourceConverterCloudDeployDeliveryPipeline() ResourceConverter {
	return ResourceConverter{
		AssetType: CloudDeployDeliveryPipelineAssetType,
		Convert:   GetCloudDeployDeliveryPipelineCaiObject,
	}
}

func GetCloudDeployDeliveryPipelineCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//clouddeploy.googleapis.com/projects/{{project}}/locations/{{region}}/deliveryPipelines/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetCloudDeployDeliveryPipelineApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: CloudDeployDeliveryPipelineAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/clouddeploy/v1/rest",
				DiscoveryName:        "DeliveryPipeline",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetCloudDeployDeliveryPipelineApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandCloudDeployDeliveryPipelineDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandCloudDeployDeliveryPipelineLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	annotationsProp, err := expandCloudDeployDeliveryPipelineAnnotations(d.Get("annotations"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("annotations"); !isEmptyValue(reflect.ValueOf(annotationsProp)) && (ok || !reflect.DeepEqual(v, annotationsProp)) {
		obj["annotations"] = annotationsProp
	}
	serialPipelineProp, err := expandCloudDeployDeliveryPipelineSerialPipeline(d.Get("serial_pipeline"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("serial_pipeline"); !isEmptyValue(reflect.ValueOf(serialPipelineProp)) && (ok || !reflect.DeepEqual(v, serialPipelineProp)) {
		obj["serialPipeline"] = serialPipelineProp
	}

	return obj, nil
}

func expandCloudDeployDeliveryPipelineDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudDeployDeliveryPipelineLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandCloudDeployDeliveryPipelineAnnotations(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandCloudDeployDeliveryPipelineSerialPipeline(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedStages, err := expandCloudDeployDeliveryPipelineSerialPipelineStages(original["stages"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStages); val.IsValid() && !isEmptyValue(val) {
		transformed["stages"] = transformedStages
	}

	return transformed, nil
}

func expandCloudDeployDeliveryPipelineSerialPipelineStages(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedTargetId, err := expandCloudDeployDeliveryPipelineSerialPipelineStagesTargetId(original["target_id"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTargetId); val.IsValid() && !isEmptyValue(val) {
			transformed["targetId"] = transformedTargetId
		}

		transformedProfiles, err := expandCloudDeployDeliveryPipelineSerialPipelineStagesProfiles(original["profiles"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedProfiles); val.IsValid() && !isEmptyValue(val) {
			transformed["profiles"] = transformedProfiles
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandCloudDeployDeliveryPipelineSerialPipelineStagesTargetId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudDeployDeliveryPipelineSerialPipelineStagesProfiles(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}