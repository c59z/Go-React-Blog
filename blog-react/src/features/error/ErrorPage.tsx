import { Button, Stack, Typography } from "@mui/material";
import { useNavigate } from "react-router-dom";
import "./ErrorPage.scss";

interface Props {
  title?: string;
  description?: string;
  showBack?: boolean;
}

const ErrorPage = ({
  title = "Something went wrong",
  description = "The page you are looking for could not be loaded.",
  showBack = true,
}: Props) => {
  const navigate = useNavigate();

  return (
    <div className="error-page">
      <Stack spacing={3} alignItems="center">
        <Typography variant="h4">{title}</Typography>

        <Typography color="text.secondary">{description}</Typography>

        <Stack direction="row" spacing={2}>
          {showBack && (
            <Button variant="outlined" onClick={() => navigate(-1)}>
              Go Back
            </Button>
          )}

          <Button variant="contained" onClick={() => navigate("/")}>
            Home
          </Button>
        </Stack>
      </Stack>
    </div>
  );
};

export default ErrorPage;
