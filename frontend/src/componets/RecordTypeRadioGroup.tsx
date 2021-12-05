import {
  FormControl,
  FormControlLabel,
  Radio,
  RadioGroup,
} from "@mui/material";
import { getListOfSupportedRecords } from "../api";
import React, { ChangeEvent } from "react";

interface RecordTypeRadioGroupProps {
  setRecordType(recordType: string): void;
  recordType: string;
}

export default function RecordTypeRadioGroup(props: RecordTypeRadioGroupProps) {
  const { recordType, setRecordType } = props;

  function updateRecordType(e: ChangeEvent<HTMLInputElement>, val: string) {
    setRecordType(val);
  }

  return (
    <FormControl component="fieldset">
      <RadioGroup
        row
        name="row-radio-buttons-group"
        onChange={updateRecordType}
        value={recordType}
      >
        {getListOfSupportedRecords().map((recordType) => (
          <FormControlLabel
            key={`key-${recordType}`}
            value={recordType}
            control={<Radio />}
            label={recordType}
            labelPlacement={"bottom"}
          />
        ))}
      </RadioGroup>
    </FormControl>
  );
}
