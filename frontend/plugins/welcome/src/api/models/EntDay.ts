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
    EntDayEdges,
    EntDayEdgesFromJSON,
    EntDayEdgesFromJSONTyped,
    EntDayEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntDay
 */
export interface EntDay {
    /**
     * Day holds the value of the "Day" field.
     * @type {string}
     * @memberof EntDay
     */
    day?: string;
    /**
     * 
     * @type {EntDayEdges}
     * @memberof EntDay
     */
    edges?: EntDayEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntDay
     */
    id?: number;
}

export function EntDayFromJSON(json: any): EntDay {
    return EntDayFromJSONTyped(json, false);
}

export function EntDayFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntDay {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'day': !exists(json, 'Day') ? undefined : json['Day'],
        'edges': !exists(json, 'edges') ? undefined : EntDayEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntDayToJSON(value?: EntDay | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'Day': value.day,
        'edges': EntDayEdgesToJSON(value.edges),
        'id': value.id,
    };
}


