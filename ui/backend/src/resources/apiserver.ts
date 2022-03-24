import { IResource, Metadata, PatchType } from '../frontend-shared';
import { getClusterApiUrl, jsonPatch, jsonRequest } from '../k8s';

export type ApiServerVersionType = 'config.openshift.io/v1';
export const ApiServerVersion: ApiServerVersionType = 'config.openshift.io/v1';

export type ApiServerKindType = 'ApiServer';
export const ApiServerKind: ApiServerKindType = 'ApiServer';

export interface NamedCertificate {
  names: string[];
  servingCertificate: { name: string };
}

export interface ApiServer extends IResource {
  apiVersion: ApiServerVersionType;
  kind: ApiServerKindType;
  metadata: Metadata;
  message?: string;
  spec?: {
    servingCerts?: { namedCertificates?: NamedCertificate[] };
  };
}

export const getApiServerConfig = (token: string) =>
  jsonRequest<ApiServer>(
    `${getClusterApiUrl()}/apis/${ApiServerVersion}/apiservers/cluster`,
    token,
  );

export const patchApiServerConfig = (token: string, patches: PatchType[]) =>
  jsonPatch<ApiServer>(
    `${getClusterApiUrl()}/apis/${ApiServerVersion}/apiservers/cluster`,
    patches,
    token,
  );
