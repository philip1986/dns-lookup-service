import { Autocomplete, TextField } from "@mui/material";
import React, { SyntheticEvent } from "react";
import { AutocompleteValue } from "@mui/base/AutocompleteUnstyled/useAutocomplete";

interface NameServerInputProps {
  setNameServer(nameServer: string[]): void;
  nameServer: string[];
}

const suggestedDnServers = [
  "default", // according to backend server config
  "8.8.8.8", // Google
  "1.1.1.1", // Cloudflare
  "9.9.9.9", // Quad9
  "77.88.8.88", // Yandex
];

export default function NameServerInput(props: NameServerInputProps) {
  const { nameServer, setNameServer } = props;

  function updateNameServer(
    e: SyntheticEvent,
    val: AutocompleteValue<string[], false, false, false>
  ) {
    if (val) setNameServer(val);
  }

  return (
    <Autocomplete
      options={suggestedDnServers}
      multiple
      freeSolo
      onChange={updateNameServer}
      value={nameServer}
      renderInput={(params) => (
        <TextField
          {...params}
          name="multiple"
          variant={"standard"}
          label="Name Servers"
        />
      )}
    />
  );
}
