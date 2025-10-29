import type { AppBarProps } from "@mui/material";
import useScrollTrigger from "@mui/material/useScrollTrigger";
import React from "react";

interface ElevationScrollProps {
  children: React.ReactElement<AppBarProps>;
}

export const ElevationScroll = ({ children }: ElevationScrollProps) => {
  const trigger = useScrollTrigger({ disableHysteresis: true, threshold: 0 });

  return React.cloneElement(children, {
    elevation: trigger ? 4 : 0,
  });
};
