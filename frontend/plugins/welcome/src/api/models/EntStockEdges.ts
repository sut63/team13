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
    EntEmployee,
    EntEmployeeFromJSON,
    EntEmployeeFromJSONTyped,
    EntEmployeeToJSON,
    EntProduct,
    EntProductFromJSON,
    EntProductFromJSONTyped,
    EntProductToJSON,
    EntTypeproduct,
    EntTypeproductFromJSON,
    EntTypeproductFromJSONTyped,
    EntTypeproductToJSON,
    EntZoneproduct,
    EntZoneproductFromJSON,
    EntZoneproductFromJSONTyped,
    EntZoneproductToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntStockEdges
 */
export interface EntStockEdges {
    /**
     * 
     * @type {EntEmployee}
     * @memberof EntStockEdges
     */
    employee?: EntEmployee;
    /**
     * 
     * @type {EntProduct}
     * @memberof EntStockEdges
     */
    product?: EntProduct;
    /**
     * 
     * @type {EntTypeproduct}
     * @memberof EntStockEdges
     */
    typeproduct?: EntTypeproduct;
    /**
     * 
     * @type {EntZoneproduct}
     * @memberof EntStockEdges
     */
    zoneproduct?: EntZoneproduct;
}

export function EntStockEdgesFromJSON(json: any): EntStockEdges {
    return EntStockEdgesFromJSONTyped(json, false);
}

export function EntStockEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntStockEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'employee': !exists(json, 'Employee') ? undefined : EntEmployeeFromJSON(json['Employee']),
        'product': !exists(json, 'Product') ? undefined : EntProductFromJSON(json['Product']),
        'typeproduct': !exists(json, 'Typeproduct') ? undefined : EntTypeproductFromJSON(json['Typeproduct']),
        'zoneproduct': !exists(json, 'Zoneproduct') ? undefined : EntZoneproductFromJSON(json['Zoneproduct']),
    };
}

export function EntStockEdgesToJSON(value?: EntStockEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'employee': EntEmployeeToJSON(value.employee),
        'product': EntProductToJSON(value.product),
        'typeproduct': EntTypeproductToJSON(value.typeproduct),
        'zoneproduct': EntZoneproductToJSON(value.zoneproduct),
    };
}


