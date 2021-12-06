import React, { useEffect, useState } from "react";
import "./App.css";
import { AppBar, Card, CardContent, Grid, Typography } from "@mui/material";
import ResultTable from "./componets/ResultTable";
import { lookupRecords, LookupResult } from "./api";
import DomainInput from "./componets/DomainInput";
import NameServerInput from "./componets/NameServerInput";
import RecordTypeRadioGroup from "./componets/RecordTypeRadioGroup";

function App() {
  const [domain, setDomain] = useState("");
  const [nameServer, setNameServer] = useState(["default"]);
  const [recordType, setRecordType] = useState("A");
  const [lookupResults, setLookupResults] = useState<LookupResult[]>([]);
  const [isLoading, setIsLoading] = useState(false);

  useEffect(() => {
    async function fetchRecords() {
      if (!domain || domain === "") {
        return;
      }
      setIsLoading(true);

      try {
        setLookupResults(await lookupRecords(domain, recordType, nameServer));
      } catch (e) {
        //TODO: show proper error message in UI
        console.error(e);
        alert(
          "Api lookup call failed. Check console or server logs for more information"
        );
      } finally {
        setIsLoading(false);
      }
    }
    fetchRecords();
  }, [domain, nameServer, recordType]);

  return (
    <Grid container rowSpacing={3} spacing={4}>
      <Grid item xs={12}>
        <AppBar position={"static"}>
          <Typography
            variant="h5"
            component="div"
            style={{ margin: "8px 8px" }}
          >
            DNS-Lookup-Tool (a.k.a Dig)
          </Typography>
        </AppBar>
      </Grid>
      <Grid item xs={1} />
      <Grid item xs={10}>
        <Card variant={"elevation"}>
          <CardContent>
            <Grid container spacing={15}>
              <Grid item xs={6}>
                <DomainInput setDomain={setDomain} />
              </Grid>
              <Grid item xs={5}>
                <NameServerInput
                  nameServer={nameServer}
                  setNameServer={setNameServer}
                />
              </Grid>
              <Grid item xs={1} />
            </Grid>
            <Grid item xs={10} style={{ marginLeft: "auto" }}>
              <RecordTypeRadioGroup
                recordType={recordType}
                setRecordType={setRecordType}
              />
            </Grid>
          </CardContent>
        </Card>
      </Grid>
      <Grid item xs={12} />
      <Grid item xs={1} />
      <Grid item xs={10}>
        <ResultTable isLoading={isLoading} data={lookupResults} />
      </Grid>
    </Grid>
  );
}

export default App;
