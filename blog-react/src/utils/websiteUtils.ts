export async function setGlobalTitle() {
  const res = await fetch("/api/website/title");
  const data = await res.json();
  document.title = data.title;
}
