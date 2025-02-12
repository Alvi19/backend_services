package util

// func IsAllowed(c *fiber.Ctx, permName string) (bool, error) {
// 	jwtUser := c.Locals("auth").(*jwt.Token)
// 	jwtClaims := jwtUser.Claims.(jwt.MapClaims)

// 	marshal, _ := json.Marshal(jwtClaims["role"])
// 	role := models.Roles{}
// 	err := json.Unmarshal(marshal, &role)
// 	if err != nil {
// 		return false, err
// 	}

// 	// for _, permission := range role.Permission {
// 	// 	if permission.Name == permName {
// 	// 		return true, nil
// 	// 	}
// 	// }

// 	return false, nil
// }
