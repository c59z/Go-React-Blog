package flag

import (
	"blog-go/global"
	"blog-go/model/appTypes"
	"blog-go/model/database"
	"blog-go/utils"
	"errors"
	"fmt"
	"os"
	"syscall"

	"github.com/gofrs/uuid"
	"golang.org/x/term"
)

// Admin creates an administrator user
func Admin() error {
	var user database.User

	// Prompt for email
	fmt.Print("Enter email: ")
	var email string
	_, err := fmt.Scanln(&email)
	if err != nil {
		return fmt.Errorf("failed to read email: %w", err)
	}
	user.Email = email

	// Get file descriptor for stdin
	fd := int(syscall.Stdin)
	// Disable echo to hide password input
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		return err
	}
	defer term.Restore(fd, oldState) // Restore terminal state

	// Prompt for password
	fmt.Print("Enter password: ")
	password, err := readPassword()
	fmt.Println()
	if err != nil {
		return err
	}

	// Prompt for password confirmation
	fmt.Print("Confirm password: ")
	rePassword, err := readPassword()
	fmt.Println()
	if err != nil {
		return err
	}

	// Check if passwords match
	if password != rePassword {
		return errors.New("passwords do not match")
	}

	// Check password length
	if len(password) < 5 || len(password) > 20 {
		return errors.New("password length should be between 5 and 20 characters")
	}

	// Populate user data
	user.UUID = uuid.Must(uuid.NewV4())
	user.Username = global.Config.Website.Name
	user.Password = utils.BcryptHash(password)
	user.RoleID = appTypes.Admin
	user.Avatar = "/image/avatar.jpg"
	user.Address = global.Config.Website.Address

	// Create user in database
	if err := global.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

// readPassword reads password input without echoing characters
func readPassword() (string, error) {
	var password string
	var buf [1]byte

	// Read characters until newline
	for {
		_, err := os.Stdin.Read(buf[:])
		if err != nil {
			return "", err
		}
		char := buf[0]

		// Stop at enter key
		if char == '\n' || char == '\r' {
			break
		}

		// Append character to password string
		password += string(char)
	}

	return password, nil
}
