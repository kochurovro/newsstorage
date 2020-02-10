package domain

type NewsStorage interface {
    GetByID(uid string) GetNewsResponse
}