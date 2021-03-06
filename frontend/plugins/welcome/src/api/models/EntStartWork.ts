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
    EntStartWorkEdges,
    EntStartWorkEdgesFromJSON,
    EntStartWorkEdgesFromJSONTyped,
    EntStartWorkEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntStartWork
 */
export interface EntStartWork {
    /**
     * StartWork holds the value of the "StartWork" field.
     * @type {string}
     * @memberof EntStartWork
     */
    startWork?: string;
    /**
     * 
     * @type {EntStartWorkEdges}
     * @memberof EntStartWork
     */
    edges?: EntStartWorkEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntStartWork
     */
    id?: number;
}

export function EntStartWorkFromJSON(json: any): EntStartWork {
    return EntStartWorkFromJSONTyped(json, false);
}

export function EntStartWorkFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntStartWork {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'startWork': !exists(json, 'StartWork') ? undefined : json['StartWork'],
        'edges': !exists(json, 'edges') ? undefined : EntStartWorkEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntStartWorkToJSON(value?: EntStartWork | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'StartWork': value.startWork,
        'edges': EntStartWorkEdgesToJSON(value.edges),
        'id': value.id,
    };
}


