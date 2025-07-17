package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"project/internal/interfaces"
	"project/internal/models"
	"project/pkg/logger"
)

type APIHandler struct {
	service interfaces.PredictionService
	log     logger.Logger
}

func NewAPIHandler(service interfaces.PredictionService, log logger.Logger) *APIHandler {
	return &APIHandler{
		service: service,
		log:     log,
	}
}

func (h *APIHandler) parseHealthData(query map[string][]string) models.HealthData {
	data := models.HealthData{
		UID: "web-client",
	}

	for key, values := range query {
		if len(values) == 0 {
			continue
		}
		value := values[0]

		switch key {
		case "uid":
			data.UID = value
		case "age":
			if age, err := strconv.Atoi(value); err == nil {
				data.Age = age
			}
		case "gender":
			if gender, err := strconv.Atoi(value); err == nil {
				data.Gender = gender
			}
		default:
			if f, err := strconv.ParseFloat(value, 64); err == nil {
				switch key {
				case "rdw":
					data.RDW = f
				case "wbc":
					data.WBC = f
				case "rbc":
					data.RBC = f
				case "hgb":
					data.HGB = f
				case "hct":
					data.HCT = f
				case "mcv":
					data.MCV = f
				case "mch":
					data.MCH = f
				case "mchc":
					data.MCHC = f
				case "plt":
					data.PLT = f
				case "neu":
					data.NEU = f
				case "eos":
					data.EOS = f
				case "bas":
					data.BAS = f
				case "lym":
					data.LYM = f
				case "mon":
					data.MON = f
				case "soe":
					data.SOE = f
				case "chol":
					data.CHOL = f
				case "glu":
					data.GLU = f
				case "hdl":
					data.HDL = f
				case "tg":
					data.TG = f
				case "crp":
					data.CRP = f
				}
			}
		}
	}

	return data
}

func (h *APIHandler) writeResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func (h *APIHandler) HandleHBA1C(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.writeResponse(w, http.StatusMethodNotAllowed, map[string]string{"error": "Method not allowed"})
		return
	}

	data := h.parseHealthData(r.URL.Query())
	response, err := h.service.Predict(r.Context(), "hba1c", data)
	if err != nil {
		h.log.Error("HBA1C prediction error", "error", err)
		h.writeResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"error":   "Prediction failed",
			"details": err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (h *APIHandler) HandleTG(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.writeResponse(w, http.StatusMethodNotAllowed, map[string]string{"error": "Method not allowed"})
		return
	}

	data := h.parseHealthData(r.URL.Query())
	response, err := h.service.Predict(r.Context(), "tg", data)
	if err != nil {
		h.log.Error("TG prediction error", "error", err)
		h.writeResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"error":   "Prediction failed",
			"details": err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (h *APIHandler) HandleHDL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.writeResponse(w, http.StatusMethodNotAllowed, map[string]string{"error": "Method not allowed"})
		return
	}

	data := h.parseHealthData(r.URL.Query())
	response, err := h.service.Predict(r.Context(), "hdl", data)
	if err != nil {
		h.log.Error("HDL prediction error", "error", err)
		h.writeResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"error":   "Prediction failed",
			"details": err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (h *APIHandler) HandleLDL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.writeResponse(w, http.StatusMethodNotAllowed, map[string]string{"error": "Method not allowed"})
		return
	}

	data := h.parseHealthData(r.URL.Query())
	response, err := h.service.Predict(r.Context(), "ldl", data)
	if err != nil {
		h.log.Error("LDL prediction error", "error", err)
		h.writeResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"error":   "Prediction failed",
			"details": err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (h *APIHandler) HandleFERR(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.writeResponse(w, http.StatusMethodNotAllowed, map[string]string{"error": "Method not allowed"})
		return
	}

	data := h.parseHealthData(r.URL.Query())
	response, err := h.service.Predict(r.Context(), "ferr", data)
	if err != nil {
		h.log.Error("FERR prediction error", "error", err)
		h.writeResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"error":   "Prediction failed",
			"details": err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (h *APIHandler) HandleLDLL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.writeResponse(w, http.StatusMethodNotAllowed, map[string]string{"error": "Method not allowed"})
		return
	}

	data := h.parseHealthData(r.URL.Query())
	response, err := h.service.Predict(r.Context(), "ldll", data)
	if err != nil {
		h.log.Error("LDLL prediction error", "error", err)
		h.writeResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"error":   "Prediction failed",
			"details": err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
