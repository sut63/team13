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
 * @interface EntGetOffWorkEdges
 */
export interface EntGetOffWorkEdges {
    /**
     * Whenendwork holds the value of the whenendwork edge.
     * @type {Array<EntEmployeeWorkingHours>}
     * @memberof EntGetOffWorkEdges
     */
    whenendwork?: Array<EntEmployeeWorkingHours>;
}

export function EntGetOffWorkEdgesFromJSON(json: any): EntGetOffWorkEdges {
    return EntGetOffWorkEdgesFromJSONTyped(json, false);
}

export function EntGetOffWorkEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntGetOffWorkEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'whenendwork': !exists(json, 'whenendwork') ? undefined : ((json['whenendwork'] as Array<any>).map(EntEmployeeWorkingHoursFromJSON)),
    };
}

export function EntGetOffWorkEdgesToJSON(value?: EntGetOffWorkEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'whenendwork': value.whenendwork === undefined ? undefined : ((value.whenendwork as Array<any>).map(EntEmployeeWorkingHoursToJSON)),
    };
}


