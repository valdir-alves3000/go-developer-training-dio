import { createUser, updateUser } from '../services/userService.js';

export function initUserForm() {
    const userForm = document.getElementById('user-form');

    userForm.addEventListener('submit', async e => {
        e.preventDefault();
        const id = document.getElementById('user-id').value;

        const userData = {
            name: document.getElementById('user-name').value,
            age: parseInt(document.getElementById('user-age').value),
            address: {
                city: document.getElementById('user-city').value,
                state: document.getElementById('user-state').value,
            },
        };

        try {
            if (id) {
                await updateUser(id, userData);
            } else {
                await createUser(userData);
            }
            window.location.reload();
        } catch (error) {
            console.error(error);
            alert('Failed to save user.');
        }
    });
}
