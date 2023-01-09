// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getAddon(args: GetAddonArgs, opts?: pulumi.InvokeOptions): Promise<GetAddonResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("heroku:addon/getAddon:getAddon", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getAddon.
 */
export interface GetAddonArgs {
    name: string;
}

/**
 * A collection of values returned by getAddon.
 */
export interface GetAddonResult {
    readonly appId: string;
    readonly configVars: string[];
    readonly id: string;
    readonly name: string;
    readonly plan: string;
    readonly providerId: string;
}
export function getAddonOutput(args: GetAddonOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAddonResult> {
    return pulumi.output(args).apply((a: any) => getAddon(a, opts))
}

/**
 * A collection of arguments for invoking getAddon.
 */
export interface GetAddonOutputArgs {
    name: pulumi.Input<string>;
}
