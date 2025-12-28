import "./index.scss";
import WebNavbar from "@/components/layout/WebNavbar";
import WebFooter from "@/components/layout/WebFooter";
import MainLayout from "../MainLayout";
import { Toolbar } from "@mui/material";

function WebLayout() {
  return (
    <>
      <WebNavbar></WebNavbar>
      <Toolbar />
      <MainLayout></MainLayout>
      <WebFooter></WebFooter>
    </>
  );
}

export default WebLayout;
