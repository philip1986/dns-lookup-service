import React from "react";

import DataTable, {
  ExpanderComponentProps,
  TableColumn,
} from "react-data-table-component";
import { LookupResult } from "../api";

interface ResultTableProps {
  data: LookupResult[];
  isLoading: boolean;
}

export type DataRow = {
  nameServer: string;
  roundTripTime: number;
  domain: string;
  ttl: number;
  recordType: string;
  value: string;
  raw: string;
};

// Table definition
const columns: TableColumn<DataRow>[] = [
  {
    name: "NameServer",
    selector: (row) => row.nameServer,
    sortable: true,
    maxWidth: "120px",
  },
  {
    name: "Domain",
    selector: (row) => row.domain,
    sortable: true,
    maxWidth: "180px",
  },
  {
    name: "Round Trip Time",
    selector: (row) => `${covertNanoIntoMilliSec(row.roundTripTime)} ms`,
    maxWidth: "120px",
  },
  {
    name: "TTL",
    selector: (row) => `${row.ttl} sec`,
    maxWidth: "120px",
  },
  {
    name: "RecordType",
    selector: (row) => row.recordType,
    sortable: true,
    maxWidth: "120px",
  },
  {
    name: "value",
    selector: (row) => row.value,
    sortable: true,
  },
];

function covertNanoIntoMilliSec(nanoSecs: number): number {
  return Math.round(nanoSecs / 1000 / 10) / 100;
}

function transformLookupResultsForDataTable(
  lookupResults: LookupResult[]
): DataRow[] {
  return lookupResults.reduce((mem, lookupResult) => {
    lookupResult.ResourceRecords.map(
      (rr) =>
        ({
          nameServer: lookupResult.NServer,
          roundTripTime: lookupResult.RoundTripTime,
          domain: rr.Domain,
          ttl: rr.TTL,
          recordType: rr.RecordType,
          value: rr.Value,
          raw: lookupResult.Raw,
        } as unknown as DataRow)
    ).forEach((i) => mem.push(i));

    return mem;
  }, [] as DataRow[]);
}

const RawView: React.FC<ExpanderComponentProps<DataRow>> = ({ data }) => {
  return <pre>{data.raw}</pre>;
};

export default function ResultTable(props: ResultTableProps) {
  return (
    <DataTable
      expandableRows
      expandableRowsComponent={RawView}
      striped
      highlightOnHover
      progressPending={props.isLoading}
      columns={columns}
      data={transformLookupResultsForDataTable(props.data)}
    />
  );
}
