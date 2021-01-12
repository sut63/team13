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
    EntPaymentchannelEdges,
    EntPaymentchannelEdgesFromJSON,
    EntPaymentchannelEdgesFromJSONTyped,
    EntPaymentchannelEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPaymentchannel
 */
export interface EntPaymentchannel {
    /**
     * Bank holds the value of the "Bank" field.
     * @type {string}
     * @memberof EntPaymentchannel
     */
    bank?: string;
    /**
     * 
     * @type {EntPaymentchannelEdges}
     * @memberof EntPaymentchannel
     */
    edges?: EntPaymentchannelEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntPaymentchannel
     */
    id?: number;
}

export function EntPaymentchannelFromJSON(json: any): EntPaymentchannel {
    return EntPaymentchannelFromJSONTyped(json, false);
}

export function EntPaymentchannelFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPaymentchannel {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'bank': !exists(json, 'Bank') ? undefined : json['Bank'],
        'edges': !exists(json, 'edges') ? undefined : EntPaymentchannelEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntPaymentchannelToJSON(value?: EntPaymentchannel | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'Bank': value.bank,
        'edges': EntPaymentchannelEdgesToJSON(value.edges),
        'id': value.id,
    };
}

