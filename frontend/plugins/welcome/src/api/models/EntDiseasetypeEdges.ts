/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntDisease,
    EntDiseaseFromJSON,
    EntDiseaseFromJSONTyped,
    EntDiseaseToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntDiseasetypeEdges
 */
export interface EntDiseasetypeEdges {
    /**
     * Disease holds the value of the disease edge.
     * @type {Array<EntDisease>}
     * @memberof EntDiseasetypeEdges
     */
    disease?: Array<EntDisease>;
}

export function EntDiseasetypeEdgesFromJSON(json: any): EntDiseasetypeEdges {
    return EntDiseasetypeEdgesFromJSONTyped(json, false);
}

export function EntDiseasetypeEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntDiseasetypeEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'disease': !exists(json, 'disease') ? undefined : ((json['disease'] as Array<any>).map(EntDiseaseFromJSON)),
    };
}

export function EntDiseasetypeEdgesToJSON(value?: EntDiseasetypeEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'disease': value.disease === undefined ? undefined : ((value.disease as Array<any>).map(EntDiseaseToJSON)),
    };
}


