package application

// type AddUserPreferencesRequest struct {
// 	UserID       int
// 	Cuisines     []int
// 	Diets        []int
// 	Intolerances []int
// }

// type AddUserPreferencesResponse struct {
// 	UserCuisines     domain.UserCuisines
// 	UserDiets        domain.UserDiets
// 	UserIntolerances domain.UserIntolerances
// }

// type AddUserPreferences func(ctx context.Context, request AddUserPreferencesRequest) (AddUserPreferencesResponse, error)

// func NewSaveUserPreferences(
// 	repo preferences.Repository,
// ) AddUserPreferences {
// 	return func(ctx context.Context, request AddUserPreferencesRequest) (AddUserPreferencesResponse, error) {

// 		var response AddUserPreferencesResponse

// 		cuisines := make(domain.UserCuisines, len(request.Cuisines))
// 		for i, cuisineID := range request.Cuisines {
// 			cuisines[i] = domain.UserCuisine{
// 				UserID:    request.UserID,
// 				CuisineID: cuisineID,
// 			}
// 		}

// 		userCuisines, err := repo.SaveUserCuisines(ctx, cuisines)
// 		if err != nil {
// 			return AddUserPreferencesResponse{}, fmt.Errorf("failed to save user cuisines: %v", err)
// 		}

// 		response.UserCuisines = userCuisines

// 		diets := make(domain.UserDiets, len(request.Diets))
// 		for i, dietID := range request.Diets {
// 			diets[i] = domain.UserDiet{
// 				UserID: request.UserID,
// 				DietID: dietID,
// 			}
// 		}

// 		userDiets, err := repo.SaveUserDiets(ctx, diets)
// 		if err != nil {
// 			return AddUserPreferencesResponse{}, fmt.Errorf("failed to save user diets: %v", err)
// 		}

// 		response.UserDiets = userDiets

// 		intolerances := make(domain.UserIntolerances, len(request.Intolerances))
// 		for i, intoleranceID := range request.Intolerances {
// 			intolerances[i] = domain.UserIntolerance{
// 				UserID:        request.UserID,
// 				IntoleranceID: intoleranceID,
// 			}
// 		}

// 		userIntolerances, err := repo.SaveUserIntolerances(ctx, intolerances)
// 		if err != nil {
// 			return AddUserPreferencesResponse{}, fmt.Errorf("failed to save user intolerances: %v", err)
// 		}
// 		response.UserIntolerances = userIntolerances

// 		return response, nil
// 	}
// }
