const form = document.getElementById("myForm");
const title = document.getElementById("title");
const content = document.getElementById("content");
const postsDiv = document.getElementById("posts");

document.cookie = "session_id=ASFALEIA_YPOLOGISTON_2024; path=/";

const fetchPosts = async () => {
  try {
    fetch("http://localhost:1234")
      .then((response) => {
        if (!response.ok) throw new Error("Failed to fetch posts");
        return response.json();
      })
      .then((posts) => {
        postsDiv.innerHTML = posts
          .map(({ title, content }) => `<p><strong>${title}</strong>: ${content}</p>`)
          .join("");
      })
      .catch(() => {
        postsDiv.innerHTML = `<p>Nothing to show yet :(</p><i>Is the server running?</i>`;
      });
  } catch (err) {
    console.error(err.message);
    elements.postsDiv.innerHTML = `<p>Nothing to show yet :(</p><i>Is the server running?</i>`;
  }
};

form.addEventListener("submit", async (event) => {
  event.preventDefault();

  if (title.value.trim() === "" || content.value.trim() === "")
    return alert("Please fill in both title and content.");

  try {
    const response = await fetch("http://localhost:1234", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ title: title.value, content: content.value }),
    });

    if (!response.ok) throw new Error("Failed to create post");

    title.value = "";
    content.value = "";

    fetchPosts();
  } catch (err) {
    console.error(err.message);
  }
});

fetchPosts();
