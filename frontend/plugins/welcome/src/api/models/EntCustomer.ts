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
    EntCustomerEdges,
    EntCustomerEdgesFromJSON,
    EntCustomerEdgesFromJSONTyped,
    EntCustomerEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntCustomer
 */
export interface EntCustomer {
    /**
     * Age holds the value of the "age" field.
     * @type {number}
     * @memberof EntCustomer
     */
    age?: number;
    /**
     * 
     * @type {EntCustomerEdges}
     * @memberof EntCustomer
     */
    edges?: EntCustomerEdges;
    /**
     * Email holds the value of the "email" field.
     * @type {string}
     * @memberof EntCustomer
     */
    email?: string;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntCustomer
     */
    id?: number;
    /**
     * Name holds the value of the "name" field.
     * @type {string}
     * @memberof EntCustomer
     */
    name?: string;
    /**
     * Password holds the value of the "password" field.
     * @type {string}
     * @memberof EntCustomer
     */
    password?: string;
}

export function EntCustomerFromJSON(json: any): EntCustomer {
    return EntCustomerFromJSONTyped(json, false);
}

export function EntCustomerFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntCustomer {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'age': !exists(json, 'age') ? undefined : json['age'],
        'edges': !exists(json, 'edges') ? undefined : EntCustomerEdgesFromJSON(json['edges']),
        'email': !exists(json, 'email') ? undefined : json['email'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'name': !exists(json, 'name') ? undefined : json['name'],
        'password': !exists(json, 'password') ? undefined : json['password'],
    };
}

export function EntCustomerToJSON(value?: EntCustomer | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'age': value.age,
        'edges': EntCustomerEdgesToJSON(value.edges),
        'email': value.email,
        'id': value.id,
        'name': value.name,
        'password': value.password,
    };
}


