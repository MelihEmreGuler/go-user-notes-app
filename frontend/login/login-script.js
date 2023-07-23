document.addEventListener("DOMContentLoaded", function() {
    const form = document.getElementById("login-form");

    form.addEventListener("submit", function(event) {
        event.preventDefault();
        const formData = new FormData(form);

        fetch("http://localhost:8080/login", {
            method: "POST",
            body: JSON.stringify({
                username_or_email: formData.get("username_or_email"),
                password: formData.get("password"),
            }),
            headers: {
                "Content-Type": "application/json",
            },
        })
            .then(response => response.json())
            .then(data => {
                // Show the response message from the backend
                alert(data.message); // Success or error message
                form.reset(); // Clear the form
            })
            .catch(error => {
                // Show error message
                alert("Error: " + error.message);
            });
    });
});
