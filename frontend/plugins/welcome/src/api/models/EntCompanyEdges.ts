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
    EntOrderproduct,
    EntOrderproductFromJSON,
    EntOrderproductFromJSONTyped,
    EntOrderproductToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntCompanyEdges
 */
export interface EntCompanyEdges {
    /**
     * Companys holds the value of the companys edge.
     * @type {Array<EntOrderproduct>}
     * @memberof EntCompanyEdges
     */
    companys?: Array<EntOrderproduct>;
}

export function EntCompanyEdgesFromJSON(json: any): EntCompanyEdges {
    return EntCompanyEdgesFromJSONTyped(json, false);
}

export function EntCompanyEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntCompanyEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'companys': !exists(json, 'companys') ? undefined : ((json['companys'] as Array<any>).map(EntOrderproductFromJSON)),
    };
}

export function EntCompanyEdgesToJSON(value?: EntCompanyEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'companys': value.companys === undefined ? undefined : ((value.companys as Array<any>).map(EntOrderproductToJSON)),
    };
}


