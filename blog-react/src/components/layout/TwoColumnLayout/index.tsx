import { Box, Stack } from "@mui/material";
import type { ReactNode } from "react";
import "./index.scss";

interface Props {
  main: ReactNode;
  aside?: ReactNode;
}

const TwoColumnLayout = ({ main, aside }: Props) => {
  return (
    <Stack direction="row" spacing={3}>
      <Box minWidth={0} flex={1}>
        {main}
      </Box>

      {aside && (
        <Box width="16rem" className="desktop-only">
          {aside}
        </Box>
      )}
    </Stack>
  );
};

export default TwoColumnLayout;
