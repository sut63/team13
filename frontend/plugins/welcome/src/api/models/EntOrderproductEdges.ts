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
    EntCompany,
    EntCompanyFromJSON,
    EntCompanyFromJSONTyped,
    EntCompanyToJSON,
    EntManager,
    EntManagerFromJSON,
    EntManagerFromJSONTyped,
    EntManagerToJSON,
    EntProduct,
    EntProductFromJSON,
    EntProductFromJSONTyped,
    EntProductToJSON,
    EntTypeproduct,
    EntTypeproductFromJSON,
    EntTypeproductFromJSONTyped,
    EntTypeproductToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntOrderproductEdges
 */
export interface EntOrderproductEdges {
    /**
     * 
     * @type {EntCompany}
     * @memberof EntOrderproductEdges
     */
    company?: EntCompany;
    /**
     * 
     * @type {EntManager}
     * @memberof EntOrderproductEdges
     */
    managers?: EntManager;
    /**
     * 
     * @type {EntProduct}
     * @memberof EntOrderproductEdges
     */
    product?: EntProduct;
    /**
     * 
     * @type {EntTypeproduct}
     * @memberof EntOrderproductEdges
     */
    typeproduct?: EntTypeproduct;
}

export function EntOrderproductEdgesFromJSON(json: any): EntOrderproductEdges {
    return EntOrderproductEdgesFromJSONTyped(json, false);
}

export function EntOrderproductEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntOrderproductEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'company': !exists(json, 'company') ? undefined : EntCompanyFromJSON(json['company']),
        'managers': !exists(json, 'managers') ? undefined : EntManagerFromJSON(json['managers']),
        'product': !exists(json, 'product') ? undefined : EntProductFromJSON(json['product']),
        'typeproduct': !exists(json, 'typeproduct') ? undefined : EntTypeproductFromJSON(json['typeproduct']),
    };
}

export function EntOrderproductEdgesToJSON(value?: EntOrderproductEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'company': EntCompanyToJSON(value.company),
        'managers': EntManagerToJSON(value.managers),
        'product': EntProductToJSON(value.product),
        'typeproduct': EntTypeproductToJSON(value.typeproduct),
    };
}


