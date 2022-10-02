const url = new URL(window.location.href);
const slug = url.searchParams.get("slug");
const date = new Date();
const copy = document.querySelector(".copy");
const yearText = document.createTextNode(`, ${date.getFullYear()}`);
copy.appendChild(yearText)
if (slug) {
	const redirectURL = `${url.origin}/u?slug=${slug}`
	const rootDir = document.getElementById("root");
	const d = document.createElement("div");
	const p = document.createElement("p");
	const a = document.createElement("a");
	a.href = redirectURL;
	a.innerText = `${redirectURL}`;
	p.innerText = "Shortened Link: ";
	d.appendChild(p);
	d.appendChild(a);
	rootDir.appendChild(d);
}