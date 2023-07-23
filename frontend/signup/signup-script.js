document.addEventListener("DOMContentLoaded", function() {
    const form = document.getElementById("signup-form");
    const consentCheckbox = document.getElementById("consent-checkbox");
    const consentPopup = document.getElementById("consent-popup");
    const closePopupBtn = document.getElementById("close-popup");
    const consentLink = document.getElementById("consent-link");
    const getStartedBtn = document.querySelector("button[type='submit']");

    form.addEventListener("submit", function(event) {
        event.preventDefault();

        const password = form.elements["password"].value;
        const confirmPassword = form.elements["confirm-password"].value;

        if (password !== confirmPassword) {
            alert("Passwords do not match. Please try again.");
            return;
        }

        if (!consentCheckbox.checked) {
            // Show the consent popup if the checkbox is not checked
            consentPopup.style.display = "block";
        } else {
            // If the checkbox is checked, proceed with form submission
            submitForm();
        }
    });

    closePopupBtn.addEventListener("click", function() {
        // Close the consent popup when the close button is clicked
        consentPopup.style.display = "none";
    });

    consentLink.addEventListener("click", function(event) {
        event.preventDefault();
        // Show the consent popup when the consent link is clicked
        consentPopup.style.display = "block";
    });

    consentCheckbox.addEventListener("change", function() {
        // Enable or disable the "Get Started" button based on checkbox state
        getStartedBtn.disabled = !this.checked;
    });

    function submitForm() {
        // Check if the consent checkbox is checked before submitting the form
        if (!consentCheckbox.checked) {
            alert("Please read and agree to the consent statement before signing up.");
            return;
        }

        const formData = new FormData(form);

        fetch("http://localhost:8080/signup", {
            method: "POST",
            body: JSON.stringify({
                username: formData.get("username"),
                email: formData.get("email"),
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

                if (data.success) {
                    // If registration is successful, redirect to login page
                    window.location.href = "../login/login.html";
                } else {
                    // If registration fails, reset the form
                    form.reset();
                }
            })
            .catch(error => {
                // Show error message
                alert("Error: " + error.message);
            });
    }
});
