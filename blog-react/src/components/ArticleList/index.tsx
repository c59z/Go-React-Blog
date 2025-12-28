import { Box, Stack, Typography } from "@mui/material";
import ArticleListItem from "../ArticleListItem";
import "./index.scss";
import ArticleSkeletonItem from "../ArticleSkeletonItem";
import { useInfiniteScroll } from "@/hook/useInfiniteScroll";
import { articleSearch, type Article } from "@/api/article";
import { useEffect, useState } from "react";
import type { Hit } from "@/api/common";

const ArticleList = () => {
  const [page, setPage] = useState<number>(1);
  const [articles, setArticles] = useState<Hit<Article>[]>([]);
  const [loading, setLoading] = useState<boolean>(false);
  const [hasMore, setHasMore] = useState<boolean>(true);

  const loadFirstPage = async () => {
    setLoading(true);
    try {
      const res = await articleSearch({
        page: 1,
        page_size: 5,
        order: "desc",
      });

      setArticles(res.data.list);
      setPage(1);
      setHasMore(res.data.list.length < res.data.total);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    loadFirstPage();
  }, []);

  const articleSkeleton = (
    <Stack spacing={3}>
      {Array.from({ length: 5 }).map((_, i) => (
        <ArticleSkeletonItem key={i} />
      ))}
    </Stack>
  );

  const fetchNextPage = async () => {
    if (loading || !hasMore) return;
    setLoading(true);
    try {
      const nextPage = page + 1;
      const res = await articleSearch({
        page: nextPage,
        page_size: 5,
        order: "desc",
      });

      setArticles((prev) => {
        const next = [...prev, ...res.data.list];
        setHasMore(next.length < res.data.total);
        return next;
      });

      setPage(nextPage);
    } finally {
      setLoading(false);
    }
  };

  const loadMoreRef = useInfiniteScroll({
    loading,
    hasMore,
    onLoadMore: fetchNextPage,
  });

  return (
    <Box className="article-list-wrapper">
      <Stack className="article-list" spacing={3}>
        {articles.map((item) => (
          <ArticleListItem
            key={item._id}
            id={item._id}
            article={item._source}
          />
        ))}
      </Stack>
      {!loading && hasMore && (
        <div ref={loadMoreRef} style={{ height: 32 }}></div>
      )}

      {!articles.length && (
        <Typography className="article-list__empty">暂无文章</Typography>
      )}
      {loading && articleSkeleton}
    </Box>
  );
};

export default ArticleList;
