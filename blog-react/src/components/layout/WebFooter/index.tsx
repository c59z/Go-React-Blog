import {
  Box,
  Typography,
  Link as MuiLink,
  GridLegacy as Grid,
} from "@mui/material";

import "./index.scss";

const WebFooter = () => {
  return (
    <Box className="footer-root">
      <Grid container justifyContent="center" wrap="nowrap">
        <Grid item xs={12} sm="auto" className="footer-grid-item">
          <div className="footer-column-style">
            <MuiLink href="#">HOME</MuiLink>
          </div>
        </Grid>

        <Grid item xs={12} sm="auto" className="footer-grid-item">
          <div className="footer-column-style">
            <MuiLink href="#">SEARCH</MuiLink>
          </div>
        </Grid>

        <Grid item xs={12} sm="auto" className="footer-grid-item">
          <div className="footer-column-style">
            <MuiLink href="#">ABOUT</MuiLink>
          </div>
        </Grid>
      </Grid>

      <Typography component="p" className="footer-copy-style">
        © 2025 c59z's Blog
      </Typography>
    </Box>
  );
};

export default WebFooter;
