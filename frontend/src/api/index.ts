// By backend service supported record types
export enum RecordType {
  "A",
  "AAAA",
  "ANY",
  "CAA",
  "CNAME",
  "MX",
  "NS",
  "PTR",
  "SOA",
  "SRV",
  "TXT",
}

type ResourceRecord = {
  Domain: string;
  TTL: number;
  RecordType: RecordType;
  Value: string;
};

export type LookupResult = {
  NServer: string;
  RoundTripTime: number;
  ResourceRecords: ResourceRecord[];
  Raw: string;
};

export async function lookupRecords(
  domain: string,
  recordType: string,
  nameServers: string[]
): Promise<LookupResult[]> {
  const nameServersStr = nameServers.join(",");
  const response = await fetch(
    `/api/v1/lookup/domain/${domain}/recordtype/${recordType}?nserver=${nameServersStr}`
  );
  return response.json();
}

export function getListOfSupportedRecords(): string[] {
  return Object.keys(RecordType).filter((k) => isNaN(Number(k)));
}
