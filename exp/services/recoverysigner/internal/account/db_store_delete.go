package account

func (s *DBStore) Delete(address string) error {
	result, err := s.DB.Exec(
		`DELETE FROM accounts
		WHERE address = $1`,
		address,
	)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}
