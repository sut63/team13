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
    EntEmployeeWorkingHours,
    EntEmployeeWorkingHoursFromJSON,
    EntEmployeeWorkingHoursFromJSONTyped,
    EntEmployeeWorkingHoursToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntShiftEdges
 */
export interface EntShiftEdges {
    /**
     * When holds the value of the when edge.
     * @type {Array<EntEmployeeWorkingHours>}
     * @memberof EntShiftEdges
     */
    when?: Array<EntEmployeeWorkingHours>;
}

export function EntShiftEdgesFromJSON(json: any): EntShiftEdges {
    return EntShiftEdgesFromJSONTyped(json, false);
}

export function EntShiftEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntShiftEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'when': !exists(json, 'when') ? undefined : ((json['when'] as Array<any>).map(EntEmployeeWorkingHoursFromJSON)),
    };
}

export function EntShiftEdgesToJSON(value?: EntShiftEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'when': value.when === undefined ? undefined : ((value.when as Array<any>).map(EntEmployeeWorkingHoursToJSON)),
    };
}


