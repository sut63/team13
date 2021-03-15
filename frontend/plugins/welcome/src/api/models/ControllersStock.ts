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
/**
 * 
 * @export
 * @interface ControllersStock
 */
export interface ControllersStock {
    /**
     * 
     * @type {number}
     * @memberof ControllersStock
     */
    amount?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersStock
     */
    employeeID?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersStock
     */
    idstock?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersStock
     */
    priceproduct?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersStock
     */
    productID?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersStock
     */
    time?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersStock
     */
    typeproductID?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersStock
     */
    zoneID?: number;
}

export function ControllersStockFromJSON(json: any): ControllersStock {
    return ControllersStockFromJSONTyped(json, false);
}

export function ControllersStockFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersStock {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'amount': !exists(json, 'amount') ? undefined : json['amount'],
        'employeeID': !exists(json, 'employeeID') ? undefined : json['employeeID'],
        'idstock': !exists(json, 'idstock') ? undefined : json['idstock'],
        'priceproduct': !exists(json, 'priceproduct') ? undefined : json['priceproduct'],
        'productID': !exists(json, 'productID') ? undefined : json['productID'],
        'time': !exists(json, 'time') ? undefined : json['time'],
        'typeproductID': !exists(json, 'typeproductID') ? undefined : json['typeproductID'],
        'zoneID': !exists(json, 'zoneID') ? undefined : json['zoneID'],
    };
}

export function ControllersStockToJSON(value?: ControllersStock | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'amount': value.amount,
        'employeeID': value.employeeID,
        'idstock': value.idstock,
        'priceproduct': value.priceproduct,
        'productID': value.productID,
        'time': value.time,
        'typeproductID': value.typeproductID,
        'zoneID': value.zoneID,
    };
}


