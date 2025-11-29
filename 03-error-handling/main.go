package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	ErrContactNotFound  = errors.New("Contact does not exist")
	ErrDuplicateContact = errors.New("Contact Already Exists")
)

type (
	ValidationError struct {
		Field   string
		Message string
	}
	FileError struct {
		Operation string
		Filename  string
		Err       error
	}
)

func (ve ValidationError) Error() string {
	return fmt.Sprintf("Invalid %s, Error : %v", ve.Field, ve.Message)
}

func (fe FileError) Error() string {
	return fmt.Sprintf("Failed %s %s, Error Message: %v", fe.Operation, fe.Filename, fe.Err)
}

type Contact struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func Validate(c Contact) error {
	if strings.TrimSpace(c.ID) == "" {
		return &ValidationError{
			Field:   "ID",
			Message: "ID cannot be empty",
		}
	}
	if strings.TrimSpace(c.Name) == "" || len(strings.TrimSpace(c.Name)) < 2 {
		return &ValidationError{
			Field:   "Name",
			Message: "Name should not be empty or less than 2 chars, enter valid name",
		}
	}
	if strings.TrimSpace(c.Email) == "" {
		return &ValidationError{
			Field:   "Email",
			Message: "Email should not be empty",
		}
	}
	if !strings.Contains(c.Email, "@") {
		return &ValidationError{
			Field:   "Email",
			Message: "Please enter Valid email with domain ",
		}
	}
	if strings.TrimSpace(c.Phone) == "" {
		return &ValidationError{
			Field:   "Phone",
			Message: "Please enter Valid Mobile Number ",
		}
	}
	return nil
}

type ContactManager struct {
	contacts map[string]Contact
	filename string
}

func NewContactManager(fileName string) *ContactManager {
	return &ContactManager{
		contacts: make(map[string]Contact),
		filename: fileName,
	}
}

func (cm *ContactManager) AddContact(contact Contact) error {

	if err := Validate(contact); err != nil {
		return fmt.Errorf("Failed to add contact %s : %w", contact.ID, err)
	}

	if _, exists := cm.contacts[contact.ID]; exists {
		return fmt.Errorf("Failed to add contact %s : %w", contact.ID, ErrDuplicateContact)
	}

	cm.contacts[contact.ID] = contact

	return nil
}

func (cm *ContactManager) GetContact(id string) (*Contact, error) {
	contact, exists := cm.contacts[id]

	if !exists {
		return nil, fmt.Errorf("Failed to fetch contact %s : %w", id, ErrContactNotFound)
	}

	return &contact, nil
}

func (cm *ContactManager) UpdateContact(id string, contact Contact) error {

	if _, exists := cm.contacts[id]; !exists {
		return fmt.Errorf("Failed to update contact %s : %w", id, ErrContactNotFound)
	}

	if err := Validate(contact); err != nil {
		return fmt.Errorf("Failed to update contact %s : %w", id, err)
	}

	contact.ID = id
	cm.contacts[id] = contact

	return nil
}

func (cm *ContactManager) DeleteContact(id string) error {
	if _, exists := cm.contacts[id]; !exists {
		return fmt.Errorf("Failed to delete contact %s : %w", id, ErrContactNotFound)
	}
	delete(cm.contacts, id)
	return nil
}

func (cm *ContactManager) ListContacts() []Contact {
	contacts := make([]Contact, len(cm.contacts))
	for _, contact := range cm.contacts {
		contacts = append(contacts, contact)
	}
	return contacts
}

func (cm *ContactManager) SaveToFile() error {
	contacts := cm.ListContacts()

	jsondata, err := json.MarshalIndent(contacts, "", " ")
	if err != nil {
		return &FileError{
			Operation: "write to",
			Filename:  cm.filename,
			Err:       err,
		}
	}

	err = os.WriteFile(cm.filename, jsondata, 0644)
	if err != nil {
		return &FileError{
			Operation: "write to",
			Filename:  cm.filename,
			Err:       err,
		}
	}

	return nil
}

func (cm *ContactManager) LoadFromFile() error {
	filedata, err := os.ReadFile(cm.filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return &FileError{
			Operation: "read from",
			Filename:  cm.filename,
			Err:       err,
		}
	}

	var contacts []Contact
	err = json.Unmarshal(filedata, &contacts)
	if err != nil {
		return &FileError{
			Operation: "read from",
			Filename:  cm.filename,
			Err:       err,
		}
	}

	for _, contact := range contacts {
		cm.contacts[contact.ID] = contact
	}

	return nil
}

func (cm *ContactManager) ImportContact(id, name, email, phone string) error {
	contact := Contact{
		ID:    id,
		Name:  name,
		Email: email,
		Phone: phone,
	}
	if err := cm.AddContact(contact); err != nil {
		return fmt.Errorf("Failed to import contact '%s': %w", name, err)
	}
	if err := cm.SaveToFile(); err != nil {
		return fmt.Errorf("Failed to save Contact '%s': %w", name, err)
	}
	return nil
}

func readInput(prompt string) string {
	fmt.Println(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func displayContact(c Contact) {
	fmt.Println(strings.Repeat("-", 10))
	fmt.Printf("ID    : %s\n", c.ID)
	fmt.Printf("Name  : %s\n", c.Name)
	fmt.Printf("Email : %s\n", c.Email)
	fmt.Printf("Phone : %s\n", c.Phone)
	fmt.Println(strings.Repeat("-", 10))
}

func handleError(err error) {
	fmt.Println("Error occured: ")
	if errors.Is(err, ErrContactNotFound) {
		fmt.Println("   The Contact you are looking for doesn't exist")
		return
	}

	if errors.Is(err, ErrDuplicateContact) {
		fmt.Println("   A Contact with this ID already exist", err)
		return
	}

	var validationErr ValidationError
	if errors.As(err, &validationErr) {
		fmt.Printf(validationErr.Error())
		return
	}

	var fileErr FileError
	if errors.As(err, &fileErr) {
		fmt.Printf(fileErr.Error())
		return
	}

	fmt.Printf(" %v\n", err)
}

func addContactMenu(cm *ContactManager) {
	fmt.Println("= Add New Contact =")
	fmt.Println("\n" + strings.Repeat("=", 50))
	id := readInput("Enter ID: ")
	name := readInput("Enter Name: ")
	email := readInput("Enter Email: ")
	phone := readInput("Enter Phone: ")
	if err := cm.ImportContact(id, name, email, phone); err != nil {
		handleError(err)
	}
	fmt.Println("Contact added successfully")
}

func viewContactMenu(cm *ContactManager) {
	fmt.Println(" Viewing Contact ")
	fmt.Println("\n" + strings.Repeat("=", 50))
	id := readInput("Enter ID to view: ")
	var contact *Contact
	var err error
	if contact, err = cm.GetContact(id); err != nil {
		handleError(err)
	}
	displayContact(*contact)
}

func updateContactMenu(cm *ContactManager) {
	fmt.Println(" Updating Contact ")
	fmt.Println("\n" + strings.Repeat("=", 50))
	id := readInput("Enter contact ID to update: ")
	existing, err := cm.GetContact(id)
	if err != nil {
		handleError(err)
	}
	fmt.Println("Current Contact:")
	displayContact(*existing)
	fmt.Println("Enter new Details, press enter to skip")
	name := readInput("Enter new Name: ")
	if name == "" {
		name = existing.Name
	}
	email := readInput("Enter new Email: ")
	if email == "" {
		email = existing.Email
	}
	phone := readInput("Enter new Phone: ")
	if phone == "" {
		phone = existing.Phone
	}
	c := Contact{ID: id, Name: name, Email: email, Phone: phone}
	if err := cm.UpdateContact(id, c); err != nil {
		handleError(err)
	}
	fmt.Println("Contact Updated SuccessFully")
}

func deleteContactMenu(cm *ContactManager) {
	fmt.Println(" Deleting Contact ")
	fmt.Println("\n" + strings.Repeat("=", 50))
	id := readInput("Enter ID to delete: ")
	existing, err := cm.GetContact(id)
	if err != nil {
		handleError(err)
	}
	fmt.Println("Current Contact:")
	displayContact(*existing)
	confirm := readInput("Are you sure to delete this contact? (yes/no): ")
	if strings.ToLower(confirm) != "yes" {
		fmt.Println("Deletion Completed")
		return
	}
	if err := cm.DeleteContact(id); err != nil {
		handleError(err)
	}
	fmt.Println("Contact Deleted Succesfully")
}

func listContactsMenu(cm *ContactManager) {
	fmt.Println(" All Contacts ")
	contacts := cm.ListContacts()
	for _, contact := range contacts {
		displayContact(contact)
	}
}

func main() {
	fmt.Println("Contact Manager - Error Handling Demo")
	fmt.Println(strings.Repeat("=", 50))

	cm := NewContactManager("contacts.json")

	fmt.Println("\nLoading contacts from file...")
	if err := cm.LoadFromFile(); err != nil {
		fmt.Println("Warning: Could not load contacts from file")
		handleError(err)
		fmt.Println("Starting with empty contact list.")
	} else {
		fmt.Printf("Loaded %d contact(s)\n", len(cm.contacts))
	}

	fmt.Println("\nDemo: Adding valid contact")
	err := cm.AddContact(Contact{
		ID:    "001",
		Name:  "John Doe",
		Email: "john@example.com",
		Phone: "555-0001",
	})
	if err != nil {
		handleError(err)
	} else {
		fmt.Println("Contact added successfully")
	}

	for {
		fmt.Println("\n" + strings.Repeat("=", 50))
		fmt.Println("Menu:")
		fmt.Println("1. Add Contact")
		fmt.Println("2. View Contact")
		fmt.Println("3. Update Contact")
		fmt.Println("4. Delete Contact")
		fmt.Println("5. List All Contacts")
		fmt.Println("6. Save and Exit")
		fmt.Println(strings.Repeat("=", 50))

		choice := readInput("Enter choice: ")

		switch choice {
		case "1":
			addContactMenu(cm)
		case "2":
			viewContactMenu(cm)
		case "3":
			updateContactMenu(cm)
		case "4":
			deleteContactMenu(cm)
		case "5":
			listContactsMenu(cm)
		case "6":
			fmt.Println("\nSaving contacts...")
			if err := cm.SaveToFile(); err != nil {
				fmt.Println("Error saving contacts:")
				handleError(err)
				confirm := readInput("Exit anyway? (yes/no): ")
				if strings.ToLower(confirm) != "yes" {
					continue
				}
			} else {
				fmt.Println("Contacts saved successfully!")
			}
			fmt.Println(" Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
