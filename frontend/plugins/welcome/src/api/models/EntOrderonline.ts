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
    EntOrderonlineEdges,
    EntOrderonlineEdgesFromJSON,
    EntOrderonlineEdgesFromJSONTyped,
    EntOrderonlineEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntOrderonline
 */
export interface EntOrderonline {
    /**
     * Addedtime holds the value of the "addedtime" field.
     * @type {string}
     * @memberof EntOrderonline
     */
    addedtime?: string;
    /**
     * 
     * @type {EntOrderonlineEdges}
     * @memberof EntOrderonline
     */
    edges?: EntOrderonlineEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntOrderonline
     */
    id?: number;
    /**
     * Stock holds the value of the "stock" field.
     * @type {number}
     * @memberof EntOrderonline
     */
    stock?: number;
}

export function EntOrderonlineFromJSON(json: any): EntOrderonline {
    return EntOrderonlineFromJSONTyped(json, false);
}

export function EntOrderonlineFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntOrderonline {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'addedtime': !exists(json, 'addedtime') ? undefined : json['addedtime'],
        'edges': !exists(json, 'edges') ? undefined : EntOrderonlineEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'stock': !exists(json, 'stock') ? undefined : json['stock'],
    };
}

export function EntOrderonlineToJSON(value?: EntOrderonline | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'addedtime': value.addedtime,
        'edges': EntOrderonlineEdgesToJSON(value.edges),
        'id': value.id,
        'stock': value.stock,
    };
}

