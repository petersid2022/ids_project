const form = document.getElementById("form");
const title = document.getElementById("title");
const content = document.getElementById("content");
const posts_div = document.getElementById("posts");

document.cookie = "session_id=ASFALEIA_YPOLOGISTON_2024; path=/";

const fetchPosts = async () => {
  try {
    const response = await fetch("http://localhost:1234");
    if (!response.ok) throw new Error("Failed to fetch posts");

    const posts = await response.json();
    if (!posts || !Array.isArray(posts)) throw new Error("Failed to deserialize JSON");

    posts_div.innerHTML = posts.map(({ title, content }) => `<p><strong>${title}</strong>: ${content}</p>`).join("");
  } catch (err) {
    console.error(err.message);
    posts_div.innerHTML = `<p>Nothing to show yet :(</p>`;
  }
};

form.addEventListener("submit", async (event) => {
  event.preventDefault();

  if (!title || !content) {
    return alert("Please fill in both the title and the content.");
  }

  try {
    const response = await fetch("http://localhost:1234", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ title: title.value, content: content.value }),
    });

    if (!response.ok) throw new Error("Failed to create post");

    title.value = title.defaultValue;
    content.value = content.defaultValue;

    fetchPosts();
  } catch (err) {
    console.error(err.message);
  }
});

fetchPosts();
