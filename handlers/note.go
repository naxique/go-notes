package handlers

import (
	"encoding/json"
	"net/http"
	"notes/models"
	"strconv"

	"github.com/gorilla/mux"
)

/*

/api/note/addnote/{userId} POST

body:
{
	"name": "...",
	"body": "..."
}

*/

func (h *Handlers) NoteAddHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(mux.Vars(r)["userId"])
	if err != nil {
		http.Error(w, "userId is not a number", http.StatusBadRequest)
		return
	}

	newNoteReq := new(models.NoteRequest)
	if err := json.NewDecoder(r.Body).Decode(newNoteReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(newNoteReq.Name) < 1 {
		http.Error(w, "Note name can't be empty", http.StatusBadRequest)
		return
	} else if len(newNoteReq.Name) > 100 {
		http.Error(w, "Note name is too long", http.StatusBadRequest)
		return
	}

	if err := h.db.CreateNewNote(userId, newNoteReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	respondWithJSON(w, http.StatusCreated, "Created")
}

/*

/api/note/getallnotes/{userId} GET

returns []Note

*/

func (h *Handlers) NoteGetAllHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(mux.Vars(r)["userId"])
	if err != nil {
		http.Error(w, "userId is not a number", http.StatusBadRequest)
		return
	}

	notes, err := h.db.GetAllNotes(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(notes) == 0 {
		http.Error(w, "No notes found for this user", http.StatusNotFound)
		return
	}

	respondWithJSON(w, http.StatusOK, notes)
}

/*

/api/note/getnote/{noteId} GET

returns Note

*/

func (h *Handlers) NoteGetHandler(w http.ResponseWriter, r *http.Request) {
	noteId, err := strconv.Atoi(mux.Vars(r)["noteId"])
	if err != nil {
		http.Error(w, "noteId is not a number", http.StatusBadRequest)
		return
	}

	note, err := h.db.GetNote(noteId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if note.ID == 0 {
		http.Error(w, "No note found with this id", http.StatusNotFound)
		return
	}

	respondWithJSON(w, http.StatusOK, note)
}

/*

/api/note/editnote/{noteId} PATCH

body:
{
	"name": "...",
	"body": "..."
}

*/

func (h *Handlers) NoteEditHandler(w http.ResponseWriter, r *http.Request) {
	noteId, err := strconv.Atoi(mux.Vars(r)["noteId"])
	if err != nil {
		http.Error(w, "noteId is not a number", http.StatusBadRequest)
		return
	}

	editNoreReq := new(models.NoteEditRequest)
	if err := json.NewDecoder(r.Body).Decode(editNoreReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(editNoreReq.Name) < 1 {
		http.Error(w, "Note name can't be empty", http.StatusBadRequest)
		return
	} else if len(editNoreReq.Name) > 100 {
		http.Error(w, "Note name is too long", http.StatusBadRequest)
		return
	}

	if err := h.db.EditNote(noteId, editNoreReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	respondWithJSON(w, http.StatusOK, "Done")
}

/*

/api/note/deletenote/{noteId} DELETE

*/

func (h *Handlers) NoteDeleteHandler(w http.ResponseWriter, r *http.Request) {
	noteId, err := strconv.Atoi(mux.Vars(r)["noteId"])
	if err != nil {
		http.Error(w, "noteId is not a number", http.StatusBadRequest)
		return
	}

	if err := h.db.DeleteNote(noteId); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	respondWithJSON(w, http.StatusOK, "Deleted")
}
