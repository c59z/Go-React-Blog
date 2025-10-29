import { Skeleton, Stack, Box } from "@mui/material";

const ArticleSkeletonItem = () => {
  return (
    <Stack direction="row" spacing={2}>
      <Skeleton variant="circular" width="6rem" height="6rem" />

      <Box flex={1}>
        <Skeleton variant="text" width="60%" height="1.5rem" />
        <Skeleton variant="text" width="40%" height="1rem" />
        <Skeleton variant="text" width="100%" />
        <Skeleton variant="text" width="90%" />
      </Box>
    </Stack>
  );
};

export default ArticleSkeletonItem;
