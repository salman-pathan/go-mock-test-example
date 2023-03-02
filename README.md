## install go-mock
`go install github.com/golang/mock/mockgen@v1.6.0`

## create mocks
`mockgen -source userRepository.go -destination userRepository_mock.go -package repo`