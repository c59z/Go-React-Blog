import { useState } from "react";
import Avatar from "@mui/material/Avatar";
import ImageOutlinedIcon from "@mui/icons-material/ImageOutlined";

const ArticleCover = ({ cover }: { cover?: string | undefined }) => {
  const [error, setError] = useState(false);

  const src = cover && !error ? cover : undefined;

  return (
    <Avatar
      variant="rounded"
      src={src}
      onError={() => setError(true)}
      sx={{
        width: "7rem",
        height: "7rem",
        bgcolor: "#f5f5f5",
        color: "#9e9e9e",
      }}
    >
      <ImageOutlinedIcon fontSize="large" />
    </Avatar>
  );
};

export default ArticleCover;
