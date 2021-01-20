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
    EntPromotionEdges,
    EntPromotionEdgesFromJSON,
    EntPromotionEdgesFromJSONTyped,
    EntPromotionEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPromotion
 */
export interface EntPromotion {
    /**
     * DurationPromotion holds the value of the "DurationPromotion" field.
     * @type {string}
     * @memberof EntPromotion
     */
    durationPromotion?: string;
    /**
     * Price holds the value of the "Price" field.
     * @type {number}
     * @memberof EntPromotion
     */
    price?: number;
    /**
     * PromotionName holds the value of the "PromotionName" field.
     * @type {string}
     * @memberof EntPromotion
     */
    promotionName?: string;
    /**
     * 
     * @type {EntPromotionEdges}
     * @memberof EntPromotion
     */
    edges?: EntPromotionEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntPromotion
     */
    id?: number;
}

export function EntPromotionFromJSON(json: any): EntPromotion {
    return EntPromotionFromJSONTyped(json, false);
}

export function EntPromotionFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPromotion {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'durationPromotion': !exists(json, 'DurationPromotion') ? undefined : json['DurationPromotion'],
        'price': !exists(json, 'Price') ? undefined : json['Price'],
        'promotionName': !exists(json, 'PromotionName') ? undefined : json['PromotionName'],
        'edges': !exists(json, 'edges') ? undefined : EntPromotionEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntPromotionToJSON(value?: EntPromotion | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'DurationPromotion': value.durationPromotion,
        'Price': value.price,
        'PromotionName': value.promotionName,
        'edges': EntPromotionEdgesToJSON(value.edges),
        'id': value.id,
    };
}


