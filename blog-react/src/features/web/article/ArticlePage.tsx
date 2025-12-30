import TwoColumnLayout from "@/components/layout/TwoColumnLayout";
import "./ArticlePage.scss";
import ArticleContent from "./components/ArticleContent";
import ArticleTOC from "./components/ArticleTOC";
import { useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import ArticleHeader from "./components/ArticleHeader";
import { Stack } from "@mui/material";
import { useArticle } from "./hook/useArticle";

export interface HeadingItem {
  id: string;
  text: string;
  level: number;
}

const ArticleDetailPage = () => {
  const [headings, setHeadings] = useState<HeadingItem[]>([]);
  const { id } = useParams();
  const { article, loading, error } = useArticle(id);
  const navigate = useNavigate();

  if (loading) {
    return <TwoColumnLayout loading={true} main={null} aside={null} />;
  }

  if (!article) {
    return null;
  }

  if (error) {
    navigate("/error", { replace: true });
    return;
  }

  return (
    <div className="article-detail-page">
      <TwoColumnLayout
        main={
          <Stack className="article-detail-body" spacing={4}>
            <ArticleHeader article={article!} />
            <ArticleContent
              content={article!.content}
              onHeadingsChange={setHeadings}
            />
          </Stack>
        }
        aside={<ArticleTOC headings={headings} />}
        loading={loading}
      />
    </div>
  );
};

export default ArticleDetailPage;
