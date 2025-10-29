import { Container } from "@mui/material";
import { Outlet } from "react-router-dom";

const MainLayout: React.FC = () => {
  return (
    <>
      <Container maxWidth="xl">
        <Outlet />
      </Container>
    </>
  );
};

export default MainLayout;
