import { Autocomplete, TextField } from "@mui/material";
import React, { ChangeEvent, SyntheticEvent, useState } from "react";
import { AutocompleteValue } from "@mui/base/AutocompleteUnstyled/useAutocomplete";

interface DomainInputProps {
  setDomain(domain: string): void;
}

export default function DomainInput(props: DomainInputProps) {
  const { setDomain } = props;
  const [partialDomain, setPartialDomain] = useState("");

  function updateDomain(
    e: SyntheticEvent,
    val: AutocompleteValue<string, false, false, true>
  ) {
    if (val) setDomain(val);
  }

  function updateDomainWhileTyping(
    e: ChangeEvent<HTMLTextAreaElement | HTMLInputElement>
  ) {
    setPartialDomain(e.currentTarget.value);
  }

  return (
    <Autocomplete
      options={[]}
      freeSolo
      onChange={updateDomain}
      onBlur={() => setDomain(partialDomain)}
      renderInput={(params) => (
        <TextField
          {...params}
          name="multiple"
          variant={"standard"}
          label="Domain"
          placeholder={"e.g.: tesla.com"}
          onChange={updateDomainWhileTyping}
        />
      )}
    />
  );
}
