import {
  AppBar,
  Toolbar,
  Box,
  IconButton,
  Button,
  Container,
} from "@mui/material";
import MenuIcon from "@mui/icons-material/Menu";
import Logo from "@/components/widgets/Logo";
import "./index.scss";
import UserAvatar from "@/components/widgets/UserAvatar";
import { ElevationScroll } from "@/components/ElevationScroll";
import { useNavigate } from "react-router-dom";

const WebNavbar = () => {
  const navigate = useNavigate();

  const handleHome = () => {
    navigate("/");
  };

  const handleBlog = () => {
    // todo go to /search
    navigate("/");
  };

  const handleAbout = () => {
    // todo go to /about
    navigate("/");
  };

  return (
    <ElevationScroll>
      <AppBar position="fixed" elevation={0} className="web-navbar">
        <Container maxWidth="xl">
          {" "}
          <Toolbar className="web-navbar-toolbar">
            {/* Logo */}
            <Box className="web-navbar-left">
              <Logo />
            </Box>

            {/* Navbar PC */}
            <Box className="web-navbar-right desktop-only">
              <Button onClick={handleHome} className="nav-item">
                Home
              </Button>
              <Button onClick={handleBlog} className="nav-item">
                Blog
              </Button>
              <Button onClick={handleAbout} className="nav-item">
                About
              </Button>
              <UserAvatar />
            </Box>

            {/* Navbar Mobile */}
            <Box className="mobile-only">
              <IconButton color="inherit">
                <MenuIcon />
              </IconButton>
            </Box>
          </Toolbar>
        </Container>
      </AppBar>
    </ElevationScroll>
  );
};

export default WebNavbar;
